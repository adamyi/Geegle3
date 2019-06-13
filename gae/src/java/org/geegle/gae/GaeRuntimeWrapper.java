package org.geegle.gae;
import java.util.concurrent.Callable;
import java.util.concurrent.Executors;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Future;
import java.util.concurrent.TimeUnit;
import java.lang.Exception;
import java.io.StringWriter;
import java.io.PrintWriter;

// idk why but Jython can't catch this exception...
// so I left this java wrapper class here
// --- adamyi
public class GaeRuntimeWrapper {
  public static Object runTask(Callable task, int timeout) {
    try {
      ExecutorService es = Executors.newSingleThreadExecutor();
      Future future = es.submit(task);
      return future.get(timeout, TimeUnit.SECONDS);
    } catch (Exception e) {
      StringWriter errors = new StringWriter();
      e.printStackTrace(new PrintWriter(errors));
      return errors.toString();
    }
  }
}
