#!/usr/bin/python3

from sys import argv
import pyautogui
import time

argc = len(argv)

if argc >= 2:
    time.sleep(float(argv[1]))

if argc == 3:
    pyautogui.press(argv[2])

if argc == 4:
    pyautogui.hotkey(argv[2], argv[3], interval=0.1)

if argc == 5:
    pyautogui.hotkey(argv[2], argv[3], argv[4], interval=0.1)
