# This is SocialAI-backend

## Launch a scalable web service in Go to handle posts and deployed to Google Cloud (Google App Engine)
## Use ElasticSearch (deployed to GCE) to provide search functgions such that users can search recent posts and list personal posts.

### run backend
`go run main.go`

### Steps to set up Google Cloud Engine
1. Enable Compute Engine
2. Set up GCE instance
3. create firewall rule
4. launch GCE instance
5. configure VSCode with GCE
    `ssh-keygen -t rsa -f ~/.ssh/gcekey -C NAME_OF_GCE_INSTANCE`
    `ls ~/.ssh/`
    `cat ~/.ssh/gcekey.pub`  get the ssh KEY_VALUE
7. In VSCode Remote Exploer add new popup window, `ssh -i ~/.ssh/gcekey YOUR_USERNAME@GCE_EXTERNAL_IP_ADDRESS`
8. choose the one under ~/.ssh/config.
9. Change SSH key and configuration permisssions
      For MacOS: `cd ~/.ssh` `chmod 600 config` `chmod 600 gcekey`
