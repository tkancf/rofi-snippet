#!/bin/bash

list="お世話になっております。\nハートビーツの山口です。
|以上、宜しくお願い致します。"

selected_text=`echo "${list[*]}" | rofi -dmenu -sep "|"`

echo -e "$selected_text" | xsel -bi

xdotool key shift+Insert
