rsync -avx ./ lgpu:/root/goctp/ --delete --progress;
ssh lgpu "export GOPATH=/mnt/ssd/projects/go && cd /root/goctp && /usr/local/go/bin/go install -v -x -a -buildmode=shared runtime sync/atomic"
ssh lgpu "export GOPATH=/mnt/ssd/projects/go && cd /root/goctp && /usr/local/go/bin/go install -v -x -work -buildmode=shared -linkshared"