# Habilita auto-cd
set -g fish_auto_cd 1

# Caminho do Go
set -x GOPATH /go
set -x GOROOT /usr/local/go
set -x PATH $PATH /go/bin

# Prompt
if status --is-interactive
    tide configure --auto
end

# Atalhos de diretório
alias ..="cd .."
alias ...="cd ../.."
alias c="clear"

# Git
alias gst="git status"
alias gaa="git add ."
alias gc="git commit -m"
alias gp="git push"
alias gl="git log --oneline --graph --decorate"

# Go
alias gor="go run ."
alias got="go test ./..."
alias gob="go build"

# Z (navegação de diretórios rápida)
z add (pwd)
