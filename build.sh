#rsync -avx ./ lgpu:/mnt/ssd/projects/go/src/github.com/geometrybase/goctp/ --delete --progress;
#ssh lgpu "export GOPATH=/mnt/ssd/projects/go && cd /mnt/ssd/projects/go/src/github.com/geometrybase/goctp/ && /usr/local/go/bin/go install -v -x -a -buildmode=shared runtime sync/atomic"
#ssh lgpu "export GOPATH=/mnt/ssd/projects/go && cd /mnt/ssd/projects/go/src/github.com/geometrybase/goctp/ && /usr/local/go/bin/go install -v -x -work -buildmode=shared -linkshared"
git add -A
git commit -m "update"
git push origin master
ssh lgpu "export GOPATH=/mnt/ssd/projects/go && /usr/local/go/bin/go get -v -u github.com/geometrybase/goctp"
