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
 - [ ] Implement auth between client <-> server to validate client is who they say they are


## Contributing
### Adding a new storage
 - [ ] Add pkg/storage/STORAGE_TYPE folder
 - [ ] Write STORAGE_TYPE implementation in that folder
   - [ ] `package STORAGE_TYPE`
   - [ ] Exported constructor of format `New(STORAGE_TYPE)Storage(config.ComponentConfig) (storage.Storage, error)`
   - [ ] Fulfill the implementation of storage.Storage (`pkg/storage/storage.go`)
   - [ ] List methods can be unordered
 - [ ] Add new storage to the `storages` slice in `pkg/storage/storage_test.go`
   - [ ] If required, create and export a function to generate a fake configuration
   - [ ] the `noOpCfg` function can be used if no configuration is required
