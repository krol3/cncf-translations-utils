## Run the github script using node

npm install @octokit/core

export GITHUB_TOKEN=xxxxx
node script.js "concepts/architecture/cri.md"

export FILE_TARGET=architecture/cri.md

## Reference
https://gist.github.com/electrocucaracha/5b274f735adb3465a26673d6c75f578f

## TO-DO
Test with go templates to render variables
fileTarget: ${FILE_TARGET}

## File handling Golang config files
- https://www.meetgor.com/golang-config-file-read/
