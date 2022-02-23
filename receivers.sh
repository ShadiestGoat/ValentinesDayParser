go build

./valentines recievers names -f | fzf \
    --preview './valentines recievers profile {}' \
    --layout=reverse\
    --bind 'enter:execute(clear && ./valentines recievers profile {})+abort'