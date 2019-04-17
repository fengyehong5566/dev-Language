#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: zhanghongyun
# Date:   2018-12-28


import time
import datetime
import logging




python学习：
	time、datetime
		时间戳：time.time()
		获取当前时间：time.localtime()
		获取当前时间并指定打印格式：time.strftime('%Y:%m:%d  %H:%M:%S')
		根据字符串转成时间对象：time.strptime('str','format')
		今天：datetime.date.today()
		昨天：datetime.date.today() - datetime.timedelta(days=-1)
		明天：datetime.date.today() - datetime.timedelta(days=1)
	








import os
	获取当前路径： os.getcwd()
	修改当前路径：os.chdir(path)
	获取指定路径下的文件名：os.listdir(path)
	创建文件夹：os.mkidr(dirname)
	创建递归文件夹：os.makedirs()
	删除目录：os.rmdir()
	递归删除空目录：os.removedirs()
	执行系统命令：os.system('ls')
	获取文件或文件夹信息：os.stat()
	重命名文件或文件夹：os.rename()

	获取系统环境变量：os.getenv()
	修改或添加环境变量：os.putenv()

	获取路径间隔符：print(os.sep)		#win:  \  , linux:  /
	获取操作系统换行符：print(repr(os.linesep))   #win:  \r\n  ,linux: \n

	获取文件名及其所在路径：os.path.split(path)  ——>  list[dirname,filename]
	合并路径和文件名：os.path.join(dirname,filename)  ——>  str:path
	获取文件的后缀：os.path.splitext(filename)  ——> list[file,postfix]
	判断是否是绝对路径：os.path.isabs(path)


	获取文件大小：os.path.getsize(filename)
	判断是否是文件、目录、连接
		os.path.isfile(filename)
		os.path.isdir(path)
		os.path.islink(filename)

	获取创建时间、修改时间、访问时间：
		os.path.getctime()
		os.path.getmtime()
		os.path.getatime()

	终止当前进程：os.exit()
	os.exit()

#文件操作：
	创建一个空文件：os.mknode()
	打开一个文件，如果不存在则创建：fp = open('filename',w)
	fp.read([size])
	fp.readline()
	fp.readlines()
	fp.write('str')    #把str内容写入文件，不会在str后加入换行符
	fp.writelines()  #一次写入多行
	fp.close()
	fp.flush()	#把缓冲区内容写入硬盘
	fp.tell()	#返回文件操作标记的当前位置，以文件的开头为原点
	fp.next()
	fp.seek()	


import sys
	#sys模块
	sys.argv	# sys.argv[0],sys.argv[1]
	sys.exit([arg])  #程序中间的退出，arg=0为正常退出
	sys.getdefaultencoding()   #获取系统当前编码
	sys.setdefaultencoding()   #设置系统默认编码
	sys.getfilesystemencoding   #获取文件使用编码
	sys.platform  #获取当前系统平台





import json
	'''
	
	dumps是将dict转化成str格式，loads是将str转化成dict格式。
	dump和load也是类似的功能，只是与文件操作结合起来了。

	'''

	json.dumps()   #将python对象编码成json字符串
	json.loads()   #将已编码的json字符串解码为python对象

	'''
	#性能比较
		for dumping
			json is faster than  simplejson;
		for loading
			simplejson is faster.

	'''
	
import  yaml
	yaml.load()
	yaml.dump() #将一个python对象生成为yaml文档




import psutil
	'''
	psutil是一个跨平台库能够轻松实现获取系统运行的进程和系统利用率（包括CPU、内存、磁盘、网络等）信息。
	它主要用来做系统监控，性能分析，进程管理。
	它实现了同等命令行工具提供的功能，如ps、top、lsof、netstat、ifconfig、who、df、kill、free、nice、ionice、iostat、iotop、uptime、pidof、tty、taskset、pmap等。
	目前支持32位和64位的Linux、Windows、OS X、FreeBSD和Sun Solaris等操作系统.

	'''
	psutil.cpu_times()  #获取cpu的完整信息
		psutil.cpu_times().user   #获取用户的cpu时间
		psutil.cpu_times().iowait	#获取IO等待时间
	psutil.cpu_count()    #获取cpu逻辑核数
	psutil.cpu_count(logical=False)   #获取cpu物理个数
	psutil.cpu_percent()	#获取cpu使用率

	psutil.virtual_memory()		#获取内存信息

	#磁盘信息[两部分]，1、磁盘的利用率； 2、IO
	psutil.disk_partitions()	#获取分区信息
	psutil.disk_io_counters()	#获取硬盘总的io数和读写信息
	psutil.disk_io_counters(perdisk=True)  #获取硬盘单个io数和读写信息

	#网络信息
	psutil.net_io_counters()	#获取网络总的IO情况
	psutil.net_io_counters(perdisk=True)	#获取网络总的IO情况

	#系统信息
	psutil.boot_time()	#获取开机时间，[时间戳]

	#进程
	psutil.pids()	#获取系统全部进程

	#查看单个进程
	p=psutil.Process(pidnum)
		p.name()	#进程名
		p.exe()		#进程的bin路径
		p.cwd()		#进程的工作目录，绝对路径
		p.create_time()		#进程创建时间
		p.num_threads()		#进程开启的线程数



import 	urllib
	'''
	http协议定义8种请求类型
		OPTIONS：返回服务器 针对 特定资源 所支持的 HTTP请求方法。也可以利用向Web服务器发送'*'的请求来测试服务器的功能性。
		HEAD：向服务器索要与GET请求相一致的响应，只不过响应体将不会被返回。
			  这一方法可以在不必传输整个响应内容的情况下，就可以获取包含在响应消息头中的元信息。 
		GET：向特定的资源发出请求。【常用】
		POST：向指定资源提交数据进行处理请求（例如提交表单或者上传文件）。数据被包含在请求体中。【常用】
			  POST请求可能会导致新的资源的创建 和/或 已有资源的修改。 
		PUT：向指定资源位置上传其最新内容
		DELETE：请求服务器删除Request-URI所标识的资源
		TRACE：回显服务器收到的请求，主要用于测试或诊断。
		CONNECT：HTTP/1.1协议中预留给能够将连接改为管道方式的代理服务器。

	'''
	#GET请求



	logging
	IPy
	dnspython
	difflib
	filecmp
	XlsxWriter
	
	rrdtool
	
	
