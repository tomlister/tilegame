import os
import subprocess

for dirpath, dnames, fnames in os.walk("./"):
	for f in fnames:
		if f.endswith(".png"):
			sp = f.split(".")
			subprocess.call(['/Users/tom.lister/go/bin/file2byteslice', '-package=assets', '-input='+dirpath+'/'+sp[0]+'.png', '-output='+sp[0]+'_go.go', '-var='+sp[0].upper()+'_go'])
            