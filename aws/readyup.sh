# TODO: Add docker-compose and json key

sudo apt update
sudo apt -y install apt-transport-https ca-certificates curl software-properties-common awscli

aws s3 cp s3://geegle/geegle-211bf7083429.json ~/
aws s3 cp s3://geegle/cluster-team-docker-compose.json ~/

echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | sudo tee -a /etc/apt/sources.list.d/google-cloud-sdk.list
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu bionic stable"
curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key --keyring /usr/share/keyrings/cloud.google.gpg add -

sudo apt update

# docker
apt-cache policy docker-ce
sudo apt install docker-ce
sudo usermod -aG docker ubuntu

# docker-compose
sudo curl -L "https://github.com/docker/compose/releases/download/1.24.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose

# gcloud
sudo apt-get install google-cloud-sdk -y

gcloud auth activate-service-account --key-file=~/geegle-211bf7083429.json
gcloud auth configure-docker

# start
sudo docker-compose -f ~/cluster-team-docker-compose.json up -d
