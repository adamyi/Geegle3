#!/bin/bash
mkdir -p $1/site/{a,b,c,d,e,f,g,h,j,i,j,k,l,m,n,o,p,q,r,s,t,u,v,y}/{0,1,2,3,4,5,6,7,8,9}/{0,1,2,3,4,5,6,7,8,9}

for file in $1/site/{a,b,c,d,e,f,g,h,j,i,j,k,l,m,n,o,p,q,r,s,t,u,v,y}/{0,1,2,3,4,5,6,7,8,9}/{0,1,2,3,4,5,6,7,8,9}/{0,1,2,3}; 
do
    echo "not the right file" > $file;
done

echo "FLAG_OR_EQUIVILANT" > $1/site/o/7/4/secret.txt
