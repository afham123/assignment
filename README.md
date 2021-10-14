# assignment

For implement of an efficient in-memory Verifiable KV-Store what we will do is that we will be creating two program, first administrator.go which only admin can access and the other public.go which everyone can access. 

## administrator.go
In this program we will first take all the news of past and store its content and its hash value(in csv) locally in our system so that the content could be kept save with us. No third party could change it. Also the locally save information could be used for the verification of the news content before serving it to any news reader.

## public.go
We will be creating a program public.go which has news data which is accessable to all. As its data is accessable to public, it has a very high risk that its content could be changed by some malicious actions. To prevent this we will checking news content before every serve. For verification we will be checking the hash values to be time efficient. If we found that the news that is being served to the user has been changed even slightly, the program will instantly restore it with the original content with the help of locally saved news content and then serve it to the user. 
    
Also the admin can add the new news to the public data with the help of userid and password.

UserId : arijit.Das
Passwowd : UITBU.Arijit
