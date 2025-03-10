import gi
gi.require_version('NM', '1.0')
from gi.repository import NM

client = NM.Client.new(None)
print ("NetworkManager version " + client.get_version())

for i in client.get_active_connections():
    print(i)
    print(i.get_default())
    for device in i.get_devices():
        print(device.get_iface())
        print(device.get_ports())
