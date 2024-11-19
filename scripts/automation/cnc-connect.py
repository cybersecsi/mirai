import telnetlib
import time

host = "localhost"
port = "11123"
user = "root"
password = "root"

wait_username_prompt = b"\033[34;1mUsername\033[33;3m: \033[0m"
wait_password_prompt = b"\033[34;1mPassword\033[33;3m: \033[0m"


def send_command(tel, command):
    """
    Sends a command over the telnet connection and reads the response.
    """
    tel.write(command.encode('utf-8') + b"\n")
    time.sleep(1)
    response = tel.read_very_eager()
    print(f"Response to '{command}':", response.decode('utf-8'))




def main():
    tel = telnetlib.Telnet(host, port, timeout=2)
    tel.write(b"\n")  # Send a newline character
    response = tel.read_until(wait_username_prompt)
    print(response.decode('utf-8'))  # Decode and print the response

    tel.write(user.encode('utf-8') + b"\n")

    response = tel.read_until(wait_password_prompt)
    print("Received:", response.decode('utf-8'))

    tel.write(password.encode('utf-8') + b"\n")
    # Optionally, read further responses
    time.sleep(1)
    final_response = tel.read_very_eager()
    print("Post-login response:", final_response.decode('utf-8'))
    # final_response = tel.read_some()
    # print("Post-login response:", final_response.decode('utf-8'))

    time.sleep(2)
    # Send the 'botcount' command
    send_command(tel, "botcount")














if __name__ == "__main__":
    main()
