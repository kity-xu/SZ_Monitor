#!usr/bin python3
# -*- coding:utf-8 -*-
import subprocess
import string

class SysCpu:

	def get_some(self):
		try:
			u = subprocess.check_output('top |head -5', shell=True).decode()
			return u
		except CalledProcessError:
			return -1

cpp = SysCpu().get_some()

print(app)