# Purpose
This file is really a TODO more than a README. If you can see this repository, you know what this project is supposed to do

### EC2 Backend
 - [ ] Accept credentials via env/file, not just configuration
 - [ ] multi-region?
 - [ ] Automatically set up IAM role? (ec2.DescribeInstances, ec2.CreateTags, ec2-instance-connect:SendSSHPublicKey, …?)
 - [ ] Pagination

## Design
 - [ ] Allow cluster-level provider configuration via handler | separate POST /cluster input from cluster.Cluster struct
 - [ ] Limit # of concurrent goroutines in job manager
 - [ ] Keep track of which entities are scaling in job manager and cancel old goroutine if newer update comes in (`singleflight` package?)
 - [ ] Implement `loop` logic to check if an instance has updated since its internal state last changed
 - [ ] Add Provisioner interface (salt, puppet, ansible, etc.) to install Phoenix agent and configure instance (instead of user data in EC2 provider)
