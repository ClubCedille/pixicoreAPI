from isc_dhcp_leases import Lease, IscDhcpLeases
from BaseHTTPServer import BaseHTTPRequestHandler, HTTPServer
import SocketServer
import cgi
import json


class Server(BaseHTTPRequestHandler):
    def _set_headers(self):
        self.send_response(200)
        self.send_header('Content-type', 'application/json')
        self.end_headers()

    def do_HEAD(self):
        self._set_headers()

    def parseDhcp(self):
        leases = IscDhcpLeases('/dhcpd/dhcpd.leases')
        data = []
        for k, v in leases.get_current().items():
            server = {"mac": k, "ip": v.ip}
            data.append(server)
        return json.dumps(data)

    def do_GET(self):
        self._set_headers()
        self.wfile.write(self.parseDhcp())


def run(server_class=HTTPServer, handler_class=Server, port=8000):
    server_address = ('', port)
    httpd = server_class(server_address, handler_class)

    print('Starting httpd on port %d...' % port)
    httpd.serve_forever()


if __name__ == "__main__":
    from sys import argv

    if len(argv) == 2:
        run(port=int(argv[1]))
    else:
        run()
