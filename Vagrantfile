# coding: utf-8
Vagrant.configure("2") do |config|

  config.vm.provider "libvirt"

  config.vm.provider :libvirt do |libvirt|
    libvirt.management_network_device = 'brPixicore'
    libvirt.driver = "kvm"
    libvirt.memory = 2048
    libvirt.cpus = 1
  end

  config.vm.define :master do |config|
    config.vm.box = "debian/stretch64"
    config.vm.hostname = 'master'
    config.vm.synced_folder ".", "/vagrant", type: "rsync",
                            rsync__exclude: ".git/"
    config.vm.network "private_network",
                      mac: '080027000011',
                      ip: '10.1.1.3',
                      libvirt__network_name: "privatePixicore",
                      autostart: "true",
                      libvirt__forward_mode: 'none',
                      libvirt__network_address: "10.1.1.0",
                      libvirt__netmask: "255.255.255.0",
                      libvirt__domain_name: "test.local",
                      libvirt__host_ip: "10.1.1.1",
                      :libvirt__dhcp_enabled => false

    config.vm.provision :shell, path: 'gateway.sh'
    config.vm.provider :libvirt do |libvirt|
      libvirt.memory = 256
    end

    config.vm.provision "docker" do |d|
      d.build_image "/vagrant/",  args:"-t pixicoreapi"
      d.run "pixicoreapi", args: "-d --network host"
    end

  end

  # Serveur 1 démarre automatiquement par pxe sur coreos
  config.vm.define :vboxNode1 do |config|
    config.vm.hostname = 'vboxNode1'
    config.vm.network "private_network",
                      mac: '080027000011',
                      libvirt__network_name: "privatePixicore",
                      autostart: "true",
                      libvirt__forward_mode: 'none',
                      libvirt__dhcp_enabled: 'false'
    config.vm.provider :libvirt do |libvirt|
      libvirt.storage :file, :size => '10G', :type => 'qcow2'
      libvirt.boot 'network'
      libvirt.boot 'hd'
      libvirt.management_network_mode = 'none'
    end
  end

  # Serveur 2 démarre automatiquement par pxe sur coreos
  config.vm.define :vboxNode2 do |config|
    config.vm.hostname = 'vboxNode1'
    config.vm.network "private_network",
                      mac: '080027000021',
                      libvirt__network_name: "privatePixicore",
                      autostart: "true",
                      libvirt__forward_mode: 'none',
                      libvirt__dhcp_enabled: 'false'
    config.vm.provider :libvirt do |libvirt|
      libvirt.storage :file, :size => '10G', :type => 'qcow2'
      libvirt.boot 'network'
      libvirt.boot 'hd'
      libvirt.management_network_mode = 'none'
      libvirt.mgmt_attach = 'false'

    end
  end

end
