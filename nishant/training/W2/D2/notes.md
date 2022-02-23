DOcker Commands

ADD
    Copies new files , directories to the image
    ADD <src> [<src>, ...] <des>

COPY
    automatically unpack compressed files
    can not take url as src

RUN 
    command after RUN is run in shell, default shell is /bin/sh -c
    RUN ['executable', 'param1', 'param2']

EXPOSE
    tells docker that container listens to the port at runtime
    doesnt bind the actual host ports


multistage builds
- multiple from statements in docker file
- each from uses a different base
- we can copy need artifacts from 1 stage to other leaving behind unnecessary stuff


continuous integration / continuous deployment

CI : 
    incremetal changes made frequently & reliably
    automated tests triggered by CI

CD : 
    every change that passes is placed into prod, making many small releases

short time to market of new product features