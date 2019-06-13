from java.util.concurrent import Callable, Executors, TimeUnit
from gaeruntime import Interpreter
from org.geegle.gae.GaeRuntimeWrapper import runTask

class ReqTask(Callable):
  def __init__(self, aeval, code):
    self.aeval = aeval
    self.code = code
  def call(self):
    self.aeval(self.code)
    return self.aeval.symtable.get('gae_rsp')

# since we don't have flask request context in java
def flattenRequest(request):
  return {"method": request.method, "cookies": request.cookies, "args": request.args, "form": request.form, "json": request.json, "host": request.host, "host_url": request.host_url, "path": request.path, "full_path": request.full_path, "url": request.url, "base_url": request.base_url, "url_root": request.url_root, "headers": request.headers, "remote_addr": request.remote_addr, "data": request.data}

def runCode(env, code, timeout):
  aeval = Interpreter(usersyms=env)
  #es = Executors.newSingleThreadExecutor()
  task = ReqTask(aeval, code)
  return runTask(task, timeout)
  #fut = es.submit(task)
  #return fut.get(timeout, TimeUnit.SECONDS)
