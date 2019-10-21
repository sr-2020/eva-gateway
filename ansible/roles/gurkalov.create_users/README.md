Create users
=========

Create and manage users with sudo and ssh for Ubuntu or Debian.

Role Variables
--------------
Generate hash password for sudo user if you need, and past them in `password` variable.

    mkpasswd --method=SHA-512
    
Also you can set variable `sudo_nopasswd: true`, and run `sudo su` without password.

Add your public key or keys in `ssh_key` variable for ssh access.

    users:
      - name: user1
        sudo: true
        sudo_nopasswd: true
        ssh_key: |
          ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQxxxxxxxxxxxxxxxxxxxA user1

      - name: user2
        sudo: true
        password: $6$Rr2XBrkWM1p$Y4gnbJrmuGwhCq0La0nxNPab2GSqR3vZVB0G3pa8BcJ224BNl1np35FLPU1FdV1b1TtpwrN4lW3OGFDKN/5FB/
        ssh_key: |
          ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQxxxxxxxxxxxxxxxxxxxB user2

Example Playbook
----------------

    - hosts: servers
      roles:
         - { role: gurkalov.create-users }

License
-------

MIT
