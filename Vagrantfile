# -*- mode: ruby -*-
# vi: set ft=ruby :

#box = "ubuntu/bionic64"
box = "bento/ubuntu-18.04"

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure("2") do |config|
  # The most common configuration options are documented and commented below.
  # For a complete reference, please see the online documentation at
  # https://docs.vagrantup.com.

  # Every Vagrant development environment requires a box. You can search for
  # boxes at https://vagrantcloud.com/search.
  config.vm.box = box


  config.vm.define "ansj1" do |server|
    server.vm.hostname = "ansj1"
    # Create a forwarded port mapping which allows access to a specific port
    # within the machine from a port on the host machine. In the example below,
    # accessing "localhost:8080" will access port 80 on the guest machine.
    # NOTE: This will enable public access to the opened port
    server.vm.network "forwarded_port", guest: 80, host: 10080
    server.vm.network "forwarded_port", guest: 3306, host: 13306

    # Create a forwarded port mapping which allows access to a specific port
    # within the machine from a port on the host machine and only allow access
    # via 127.0.0.1 to disable public access
    # server.vm.network "forwarded_port", guest: 80, host: 8080, host_ip: "127.0.0.1"

    # Create a private network, which allows host-only access to the machine
    # using a specific IP.
    # server.vm.network "private_network", ip: "192.168.33.10"
    server.vm.network "private_network", type: "dhcp"

    # Create a public network, which generally matched to bridged network.
    # Bridged networks make the machine appear as another physical device on
    # your network.
    # server.vm.network "public_network"

    server.vm.synced_folder "./share", "/vagrant_data"

    server.vm.provider "virtualbox" do |vb|
      vb.name = "ansj1"
      vb.gui = false
      vb.memory = "1024"
    end

    # Enable provisioning with a shell script. Additional provisioners such as
    # Puppet, Chef, Ansible, Salt, and Docker are also available. Please see the
    # documentation for more information about their specific syntax and use.
    # config.vm.provision "shell", inline: <<-SHELL
    #   apt-get update
    #   apt-get install -y apache2
    # SHELL

    server.vm.provision "shell", inline: <<-SHELL
      set -e
      export DEBIAN_FRONTEND=noninteractive
      apt-get update
      apt-get install -y --no-install-recommends ansible git
    SHELL
  end

  config.vm.define "ansj2" do |server|
    server.vm.hostname = "ansj2"
    # Create a forwarded port mapping which allows access to a specific port
    # within the machine from a port on the host machine. In the example below,
    # accessing "localhost:8080" will access port 80 on the guest machine.
    # NOTE: This will enable public access to the opened port
    server.vm.network "forwarded_port", guest: 80, host: 20080
    server.vm.network "forwarded_port", guest: 3306, host: 23306

    # Create a forwarded port mapping which allows access to a specific port
    # within the machine from a port on the host machine and only allow access
    # via 127.0.0.1 to disable public access
    # server.vm.network "forwarded_port", guest: 80, host: 8080, host_ip: "127.0.0.1"

    # Create a private network, which allows host-only access to the machine
    # using a specific IP.
    # server.vm.network "private_network", ip: "192.168.33.10"
    server.vm.network "private_network", type: "dhcp"

    # Create a public network, which generally matched to bridged network.
    # Bridged networks make the machine appear as another physical device on
    # your network.
    # server.vm.network "public_network"

    server.vm.synced_folder "./share", "/vagrant_data"

    server.vm.provider "virtualbox" do |vb|
      vb.name = "ansj2"
      vb.gui = false
      vb.memory = "1024"
    end

    # Enable provisioning with a shell script. Additional provisioners such as
    # Puppet, Chef, Ansible, Salt, and Docker are also available. Please see the
    # documentation for more information about their specific syntax and use.
    # server.vm.provision "shell", inline: <<-SHELL
    #   apt-get update
    #   apt-get install -y apache2
    # SHELL

    server.vm.provision "shell", inline: <<-SHELL
      set -e
      export DEBIAN_FRONTEND=noninteractive
      apt-get update
      apt-get install -y --no-install-recommends ansible git
    SHELL
  end

  config.vm.define "ansj3" do |server|
    server.vm.hostname = "ansj3"
    # Create a forwarded port mapping which allows access to a specific port
    # within the machine from a port on the host machine. In the example below,
    # accessing "localhost:8080" will access port 80 on the guest machine.
    # NOTE: This will enable public access to the opened port
    server.vm.network "forwarded_port", guest: 80, host: 30080
    server.vm.network "forwarded_port", guest: 3306, host: 33306

    # Create a forwarded port mapping which allows access to a specific port
    # within the machine from a port on the host machine and only allow access
    # via 127.0.0.1 to disable public access
    # server.vm.network "forwarded_port", guest: 80, host: 8080, host_ip: "127.0.0.1"

    # Create a private network, which allows host-only access to the machine
    # using a specific IP.
    # server.vm.network "private_network", ip: "192.168.33.10"
    server.vm.network "private_network", type: "dhcp"

    # Create a public network, which generally matched to bridged network.
    # Bridged networks make the machine appear as another physical device on
    # your network.
    # server.vm.network "public_network"

    server.vm.synced_folder "./share", "/vagrant_data"

    server.vm.provider "virtualbox" do |vb|
      vb.name = "ansj3"
      vb.gui = false
      vb.memory = "1024"
    end

    # Enable provisioning with a shell script. Additional provisioners such as
    # Puppet, Chef, Ansible, Salt, and Docker are also available. Please see the
    # documentation for more information about their specific syntax and use.
    # server.vm.provision "shell", inline: <<-SHELL
    #   apt-get update
    #   apt-get install -y apache2
    # SHELL

    server.vm.provision "shell", inline: <<-SHELL
      set -e
      export DEBIAN_FRONTEND=noninteractive
      apt-get update
      apt-get install -y --no-install-recommends ansible git
    SHELL
  end
end
