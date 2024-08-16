cd testutil/vector/config

# delete pem file
rm *.pem

# Create CA private key and self-signed certificate
# adding -nodes to not encrypt the private key
openssl req -x509 -newkey rsa:4096 -nodes -days 365 -keyout ca-key.pem -out ca-cert.pem -subj "/C=US/ST=NA/L=NY/O=LAGRANGE/OU=SIGNER/CN=*.signer.ca/emailAddress=signer@ca.com"

# Create Server private key and CSR
# adding -nodes to not encrypt the private key
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=US/ST=NA/L=NY/O=LAGRANGE/OU=SIGNER/CN=*.signer.srv/emailAddress=signer@srv.com"

# Sign the Server Certificate Request (CSR)
openssl x509 -req -in server-req.pem -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile ../ext.conf

# Verify certificate
openssl verify -CAfile ca-cert.pem server-cert.pem

# Generate client's private key and certificate signing request
openssl req -newkey rsa:4096 -nodes -keyout client-key.pem -out client-req.pem -subj "/C=US/ST=NA/L=NY/O=LAGRANGE/OU=SIGNER/CN=*.signer.client/emailAddress=signer@client.com"

# Sign the Client Certificate Request (CSR)
openssl x509 -req -in client-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-cert.pem -extfile ../ext.conf