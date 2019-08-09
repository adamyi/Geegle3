javac /geegle/gae/java/org/geegle/gae/*.java
mkdir /src/WEB-INF/lib/
cp /usr/local/jython/jython.jar /src/WEB-INF/lib/jython.jar
mkdir /src/WEB-INF/lib-python/
mv /python_requirements/*/* /src/WEB-INF/lib-python/
\rm -rf /python_requirements
catalina.sh run
