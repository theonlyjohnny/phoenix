### EC2 Backend
 - [ ] Accept creds via env/file, not just config
 - [ ] multi-region?
 - [ ] Automatically set up IAM role? (ec2.describeinstances, 

## Design
 - [ ] Move manager calls into storage so each handler doesn't have to implement? <-- Then does each storage have to implement? annoying if eventually wanna make plugins
