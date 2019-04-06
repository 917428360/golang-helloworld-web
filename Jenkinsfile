String buildShell = "${env.buildShell}"
String targetHosts = "${env.targetHosts}"
String targetDir = "${env.targetDir}"
String serviceName = "${env.serviceName}"
String user = "${env.user}"

def jarName

node("master"){
    stage("checkout"){
        checkout scm
    }
    
    stage("build"){
        sh "${buildShell} "
        
        sh """ mkdir -p /srv/salt/${serviceName} 
               tar zcf ${serviceName}.tar.gz main static service.sh 
               rm -fr /srv/salt/${serviceName}/*
               mv ${serviceName}.tar.gz /srv/salt/${serviceName} 
           """
    }
    
    stage("deploy"){
        sh " salt ${targetHosts} cmd.run ' rm -fr  ${targetDir}/* '"
        sh " salt ${targetHosts} cp.get_file salt://${serviceName}/${serviceName}.tar.gz  ${targetDir}/${serviceName}.tar.gz mkdirs=True"
        sh " salt ${targetHosts} cmd.run 'chown ${user}:${user} ${targetDir} -R '"
        sh " salt ${targetHosts} cmd.run 'su - ${user} -c \"cd ${targetDir} &&  sh service.sh stop\" ' "
        sh " salt ${targetHosts} cmd.run 'su - ${user} -c \"cd ${targetDir} &&  sh service.sh start ${targetDir}\" ' "
    }


}
