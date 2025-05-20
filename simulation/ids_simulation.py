import os
import time
import random
import shutil
import subprocess
from threading import Thread
from mininet.net import Mininet
from mininet.node import OVSSwitch
from mininet.link import TCLink
from mininet.log import setLogLevel, info
import psycopg2
from psycopg2 import sql

def clear_logs(*dirs):
    for d in dirs:
        if os.path.exists(d):
            shutil.rmtree(d)
        os.makedirs(d)

def insert_alert_to_db(log_source, alert_text):
    try:
        conn = psycopg2.connect(
            dbname="snortdb",
            user="snortuser",
            password="123456",
            host="localhost"
        )
        cur = conn.cursor()
        insert_query = sql.SQL(
            "INSERT INTO snort_alerts (log_source, alert_text) VALUES (%s, %s);"
        )
        cur.execute(insert_query, (log_source, alert_text))
        conn.commit()
        cur.close()
        conn.close()
    except Exception as e:
        print(f"[DB ERROR] {e}")

def setup_snort(host, log_dir):
    iface = f"{host.name}-eth0"
    host.cmd(f'mkdir -p {log_dir}')

    rules_path = '/etc/snort/rules/local.rules'
    rule_lines = [
        'alert icmp any any -> any any (msg:"ICMP Packet Detected"; sid:1000001; rev:1;)',
        'alert tcp any any -> any any (flags:S; msg:"SYN Packet Detected"; sid:1000002; rev:1;)',
        'alert udp any any -> any any (msg:"UDP Packet Detected"; sid:1000003; rev:1;)',
        'alert tcp any any -> any any (flags:S; msg:"Port Scan Attempt"; detection_filter:track by_src, count 3, seconds 10; sid:1000004; rev:2;)'
    ]

    if os.path.exists(rules_path):
        existing = open(rules_path).read()
    else:
        existing = ""
        os.makedirs(os.path.dirname(rules_path), exist_ok=True)

    with open(rules_path, 'a') as f:
        for rule in rule_lines:
            if rule not in existing:
                f.write(rule + "\n")

    snort_cmd = [
        "snort",
        "-A", "fast",
        "-i", iface,
        "-c", "/etc/snort/snort.conf",
        "-l", log_dir
    ]

    print(f"[+] Launching Snort on {host.name} with: {' '.join(snort_cmd)}")
    return host.popen(snort_cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE)

def read_snort_logs(log_dir, last_seen):
    alert_file = os.path.join(log_dir, 'alert')
    if not os.path.exists(alert_file):
        return "", last_seen

    with open(alert_file, 'r') as f:
        lines = f.readlines()
        new_lines = lines[last_seen:]
        last_seen = len(lines)

    return ''.join(new_lines), last_seen

def icmp_flood(attacker, target):
    print(f"[*] ICMP Flood: {attacker.name} -> {target.name}")
    attacker.cmd(f'ping -i 0.2 -c 10 {target.IP()} > /dev/null &')

def syn_flood(attacker, target):
    print(f"[*] SYN Flood: {attacker.name} -> {target.name}")
    attacker.cmd(f'hping3 -S -p 80 -i u10000 -c 10 {target.IP()} > /dev/null 2>&1 &')

def udp_flood(attacker, target):
    print(f"[*] UDP Flood: {attacker.name} -> {target.name}")
    attacker.cmd(f'hping3 --udp -p 53 -i u10000 -c 10 {target.IP()} > /dev/null 2>&1 &')

def port_scan(attacker, target):
    print(f"[*] Port Scan: {attacker.name} -> {target.name}")
    attacker.cmd(f'nmap -Pn -sS -T4 {target.IP()} > /dev/null 2>&1 &')

def attack_loop(attacker, targets):
    attacks = [icmp_flood, syn_flood, udp_flood, port_scan]
    while True:
        target = random.choice(targets)
        attack = random.choice(attacks)
        attack(attacker, target)
        time.sleep(10)

def alert_printer(log_dirs):
    last_seen = {d: 0 for d in log_dirs}
    while True:
        for log_dir in log_dirs:
            logs, last_seen[log_dir] = read_snort_logs(log_dir, last_seen[log_dir])
            if logs:
                for line in logs.strip().splitlines():
                    print(f"[!] Alert from {log_dir}: {line}")
                    insert_alert_to_db(log_dir, line)
        time.sleep(5)

def run_simulation():
    os.system('mn -c')
    log_dir_h1 = '/tmp/snort_logs_h1'
    log_dir_h2 = '/tmp/snort_logs_h2'
    clear_logs(log_dir_h1, log_dir_h2)

    net = Mininet(controller=None, switch=OVSSwitch, link=TCLink)

    info("*** Creating nodes\n")
    h1 = net.addHost('h1', ip='10.0.0.1/24')
    h2 = net.addHost('h2', ip='10.0.0.2/24')
    attacker = net.addHost('attacker', ip='10.0.0.99/24')
    s1 = net.addSwitch('s1')

    info("*** Creating links\n")
    net.addLink(h1, s1)
    net.addLink(h2, s1)
    net.addLink(attacker, s1)

    info("*** Starting network\n")
    net.start()

    setup_snort(h1, log_dir_h1)
    setup_snort(h2, log_dir_h2)

    print("[*] Waiting for Snort to initialize...")
    time.sleep(5)

    Thread(target=attack_loop, args=(attacker, [h1, h2]), daemon=True).start()
    Thread(target=alert_printer, args=([log_dir_h1, log_dir_h2],), daemon=True).start()

    print("[*] Snort IDS simulation is running. Press Ctrl+C to stop.")
    while True:
        time.sleep(60)

if __name__ == '__main__':
    setLogLevel('info')
    run_simulation()