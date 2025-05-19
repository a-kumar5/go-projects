# go-projects

`export PATH=$PATH:$(go env GOPATH)/bin`


gcloud iam service-accounts add-iam-policy-binding "github-wif@able-current-451504-j6.iam.gserviceaccount.com" \
  --project="tt-dev-001" \
  --role="roles/iam.workloadIdentityUser" \
  --member="principalSet://iam.googleapis.com/projects/517338717512/locations/global/workloadIdentityPools/github/attribute.repository/a-kumar5/go-projects"