#!/bin/bash

cd "/Users/tomas.kocman/Desktop/lenslocked"

echo "==== Releasing lenslocked.com ===="
echo " Deleting the local binary if it exists (so it isn't uploaded)..."
rm lenslocked.com
echo " Done!"

echo " Deleting existing code..."
ssh root@139.59.141.98 "rm -rf /root/go/src/lenslocked.com"
echo " Code deleted successfully!"

echo " Uploading code..."
rsync -avr --exclude '.git/*' --exclude 'tmp/*' --exclude 'images/*' ./ root@139.59.141.98:/root/go/src/lenslocked.com/
echo "  Code uploaded successfully!"

echo " Go getting deps..."
ssh root@139.59.141.98 "cd /root/go/src/lenslocked.com/; /usr/local/go/bin/go mod download && /usr/local/go/bin/go mod verify"

echo " Building the code on remote server..."
ssh root@139.59.141.98 "cd /root/go/src/lenslocked.com/; /usr/local/go/bin/go build -o /root/app/server -i ./cmd/lenslocked/*"
echo " Code built successfully!"

echo " Moving assets..."
ssh root@139.59.141.98 "cd /root/app; cp -R /root/go/src/lenslocked.com/assets ."
echo " Assets moved successfully!"

echo " Moving views..."
ssh root@139.59.141.98 "cd /root/app; cp -R /root/go/src/lenslocked.com/views ."
echo " Views moved successfully!"

echo " Moving Caddyfile..."
ssh root@139.59.141.98 "cd /root/app; cp /root/go/src/lenslocked.com/Caddyfile ."
echo " Caddyfile moved successfully!"

echo " Restarting the server..."
ssh root@139.59.141.98 "service lenslocked.com restart"
echo " Server restarted successfully!"

echo " Restarting Caddy server..."
ssh root@139.59.141.98 "service caddy restart"
echo " Caddy restarted successfully!"

echo "==== Done releasing lenslocked.com ===="
