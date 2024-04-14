import subprocess

servers = [
    {
        "name": "Server ONE:2221",
        "port": "2221"
    },
    {
        "name": "Server TWO:2222",
        "port": "2222"
    },
    {
        "name": "Server THREE:2223",
        "port": "2223"
    },
    {
        "name": "Server FOUR:2224",
        "port": "2224"
    },
    {
        "name": "Server FIVE:2225",
        "port": "2225"
    },
    {
        "name": "Server SIX:2226",
        "port": "2226"
    },
    {
        "name": "Server SEVEN:2227",
        "port": "2227"
    },
    {
        "name": "Server EIGHT:2228",
        "port": "2228"
    },
    {
        "name": "Server NINE:2229",
        "port": "2229"
    },
    {
        "name": "Server TEN:2230",
        "port": "2230"
    },
]

if __name__ == "__main__":
    for server in servers:
        port_mapping = f"{server['port']}:8001"
        command = [
            "docker", "run", "-d",
            "-e", "PRACTICE_APP_PORT=8001",
            "-e", f"CURRENT_APP={server['name']}",
            "-p", port_mapping,
            "haproxy-practice-app"
        ]
        
        try:
            subprocess.run(command, check=True)
        except subprocess.CalledProcessError as e:
            print(f"Error executing command for {server['name']}: {e}")

