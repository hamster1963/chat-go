# 获取git tag版本号
GitTag=$(git tag --sort=version:refname | tail -n 1)
# 获取源码最近一次 git commit log，包含 commit sha 值，以及 commit message
GitCommitLog=$(git log --pretty=oneline -n 1)
# 获取当前时间
BuildTime=$(date +'%Y.%m.%d.%H:%M:%S')
# 获取Go的版本
BuildGoVersion=$(go version)

# 打印
echo "GitTag: ${GitTag}"
echo "GitCommitLog: ${GitCommitLog}"
echo "BuildTime: ${BuildTime}"
echo "BuildGoVersion: ${BuildGoVersion}"

# 将以上变量序列化至 LDFlags 变量中
LDFlags=" \
    -X 'push/utility/bin_utils.GitTag=${GitTag}' \
    -X 'push/utility/bin_utils.GitCommitLog=${GitCommitLog}' \
    -X 'push/utility/bin_utils.BuildTime=${BuildTime}' \
    -X 'push/utility/bin_utils.BuildGoVersion=${BuildGoVersion}' \
"

go build -ldflags "$LDFlags" -o service