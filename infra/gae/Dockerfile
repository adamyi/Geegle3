FROM tomcat:9.0.21-jdk8

ENV JYTHON_VERSION="2.7.1" \
    JYTHON_SHASUM="392119a4c89fa1b234225d83775e38dbd149989f"

RUN curl -fSL -o jython_installer.jar "http://search.maven.org/remotecontent?filepath=org/python/jython-installer/${JYTHON_VERSION}/jython-installer-${JYTHON_VERSION}.jar" \
 && echo "$JYTHON_SHASUM *jython_installer.jar" | sha1sum -c - \
 && java -jar jython_installer.jar -s -d /usr/local/jython \
 && rm jython_installer.jar \
 && ln -s /usr/local/jython/bin/jython /usr/local/jython/bin/python \
 && ln -s /usr/local/jython/bin/jython /usr/local/jython/bin/python2

ENV PATH="/usr/local/jython/bin:$PATH"

# separating this makes incremental building fast
COPY src/requirements.txt /
RUN pip install -r /requirements.txt

RUN mkdir /geegle

COPY src /geegle/gae

RUN javac /geegle/gae/java/org/geegle/gae/*.java

RUN mkdir /geegle/gae/WEB-INF/lib \
 && cp /usr/local/jython/jython.jar /geegle/gae/WEB-INF/lib

WORKDIR /geegle/gae

RUN rm -rf /usr/local/tomcat/webapps/*
COPY tomcat/ROOT.xml /usr/local/tomcat/conf/Catalina/localhost/ROOT.xml
COPY tomcat/server.xml /usr/local/tomcat/conf/server.xml

ENV JYTHONPATH "/geegle/gae/java"

