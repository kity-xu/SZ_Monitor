import time
import sys
sys.path.append(r'monitor.im/local/utils')
import tool

def redisProcess(hcmd, cmd):
	#redis = redis_server(hcmd, cmd, 'xulang')
	data = {}
	head = tool.getHead(hcmd)
	info = tool.getInfo(cmd)
	
	for i in range(len(head)):
		if len(head)-1 == i:
			data[head[i]] = info[i] + ' ' + info[i+1]
			break
		data[head[i]] = info[i]
	return (str(data))
	
	#return ("{0},{1}".format(str(data),id(redis)))

#if __name__ == '__main__':
print(redisProcess('ps -aux |head -1', 'ps -aux | grep redis-server'))