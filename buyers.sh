go build
./valentines buyers names -f | fzf \
    --preview './valentines buyers profile {}' \
    --layout=reverse\
    --bind 'enter:execute(clear && ./valentines buyers profile {})+abort'