#!/usr/bin/python3

from sys import argv
import pyautogui

argc = len(argv)

if argc == 2:
    pyautogui.press(argv[1])

if argc == 3:
    pyautogui.hotkey(argv[1], argv[2], interval=0.1)

if argc == 4:
    pyautogui.hotkey(argv[1], argv[2], argv[3], interval=0.1)
