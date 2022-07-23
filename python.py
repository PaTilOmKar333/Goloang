from os import stat


def palandrome(a):
    mid=len(a)/2
    start=0
    last=len(a)-1
    flag=0

    while(start < last) :
        if (a[start]==a[last]):
            start+=1
            last-+1
        else:
            flag=1
            break

    if(flag==0):
        print("flag:",flag)
        print("this is palandrom string")
    else:
        print("flag:",flag)
        print("this is not palandrom string")



str="aaaaaaaaa"
palandrome(str)



