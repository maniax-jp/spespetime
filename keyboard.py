#!/usr/bin/python3

"""
引数に従ってキーボード操作を行う
引数1: 待機時間(秒)
引数2: キー
引数3: キー
引数4: キー

例1: 0.1秒待ってEnterキーを押す
$ python3 keyboard.py 0.1 Enter

例2: 0.1秒待ってCtrl + Shift + F1を押す
$ python3 keyboard.py 0.1 ctrl shift f1
"""

from sys import argv
import time

import pyautogui

ARGC = len(argv)

if ARGC >= 2:
    time.sleep(float(argv[1]))

if ARGC == 3:
    pyautogui.press(argv[2])

if ARGC == 4:
    pyautogui.hotkey(argv[2], argv[3], interval=0.1)

if ARGC == 5:
    pyautogui.hotkey(argv[2], argv[3], argv[4], interval=0.1)
