## Golang code samples.
Generate executables
```
go build -o ./app1/app1 ./app1/main.go && go build -o ./app2/app2 ./app2/main.go
```
Copy services to `/lib/systemd`
```
cp ./services/app1.service /lib/systemd/system/ 
cp ./services/app2.service /lib/systemd/system/
```
Copy nginx conf to `/etc/nginx` and create symlink to enabled sites
```
sudo cp nginx.conf /etc/nginx/sites-available/exampleapp
sudo ln -s /etc/nginx/sites-available/exampleapp /etc/nginx/sites-enabled/exampleapp
```
Add to exampleapp domain to `/etc/hosts`
```
...
127.0.0.1       exampleapp
```
Run services with systemctl
```
systemctl start app1
systemctl start app2
```
Test
```
curl -v http://exampleapp/app2/ 
curl -v http://exampleapp/app1/
```
