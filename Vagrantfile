Vagrant.configure("2") do |config|
  config.vm.box = "debian/stretch64"

  config.vm.provider :virtualbox do |vb|
    vb.linked_clone = true
    vb.memory = 256
    vb.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
    vb.customize ['modifyvm', :id, '--hostonlyadapter1', "vboxnet5"]
    vb.customize ['modifyvm', :id, '--cableconnected2', 'on']
  end

  config.vm.define :master do |config|
    config.vm.hostname = 'master'
    config.vm.network "private_network", mac: '080027000011', ip: '10.1.1.3', name: 'vboxnet5'
    config.vm.provision :shell, path: 'gateway.sh'
    config.vm.provision "docker" do |d|
      d.build_image "/vagrant/",  args:"-t pixicoreapi" 
      d.run "pixicoreapi", args: "-d --network host -it --name pixicoreapi -v /var/lib/dhcp:/dhcpd"
    end
    config.vm.provider :virtualbox do |vb, config|
      vb.memory = 2048
      vb.customize ['modifyvm', :id, '--macaddress1', '080027000012']
      vb.customize ['modifyvm', :id, '--natnet1', '10.0.2.0/24']
      vb.customize ['modifyvm', :id, '--nictype4', 'virtio']
    end
  end

  # Serveur 1 démarre automatiquement par pxe sur coreos
  config.vm.define :vboxNode1 do |config|
    config.vm.hostname = 'vboxNode1'
    config.vm.provider :virtualbox do |vb, config|
      vb.memory = 2048
      vb.customize ['modifyvm', :id, '--nic1', "hostonly"]
      vb.customize ['modifyvm', :id, '--macaddress1', '080027000021']
      vb.customize ['modifyvm', :id, '--natnet1', '10.0.2.0/24']
      vb.customize ['modifyvm', :id, '--boot1', 'net']
      vb.customize ['modifyvm', :id, '--biospxedebug', 'on']
      vb.customize ['modifyvm', :id, '--nicbootprio2', '1']
      vb.customize ['modifyvm', :id, "--nictype2", '82540EM']
    end
  end

  # Serveur 2 démarre automatiquement par pxe sur coreos
  config.vm.define :vboxNode2 do |config|
    config.vm.hostname = 'vboxNode2'
    config.vm.provider :virtualbox do |vb, config|
      vb.memory = 2048
      vb.customize ['modifyvm', :id, '--nic1', "hostonly"]
      vb.customize ['modifyvm', :id, '--macaddress1', '080027000031']
      vb.customize ['modifyvm', :id, '--natnet1', '10.0.2.0/24']
      vb.customize ['modifyvm', :id, '--boot1', 'net']
      vb.customize ['modifyvm', :id, '--biospxedebug', 'on']
      vb.customize ['modifyvm', :id, '--nicbootprio2', '1']
      vb.customize ['modifyvm', :id, "--nictype2", '82540EM']
    end
  end
  
end 