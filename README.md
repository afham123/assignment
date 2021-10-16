# assignment

For implement of an efficient in-memory Verifiable KV-Store what we will do is that we will be creating two program, first administrator.go which only admin can access and the other public.go which everyone can access. 

## administrator.go
In this program we will first take all the news of past and store its content and its hash value(in csv) locally in our system so that the content could be kept save with us. No third party could change it. Also the locally save information could be used for the verification of the news content before serving it to any news reader. 
### details about the program.
We have firstly taken the news and stored it in hash map and after that saved the news in csv format locally in our system. Functions used in administration.go and their usage are:
* Put(key,val) : Take the key, value pair from the user and store it in the hash map (keyValue)
* save() : Function to save the content locally from hash map to csv. 
* Delete(key) : Function to delete the content(value and hash).
* Get(key) :  Function to retrieve news contnet and its Hash Value stored in a hash map.
* Put(key, value) : Function that takes key, value pair as an argument and store it in key,value and its hash256 in keyValue hash map.
* readString(buffer) : Function to read string.
* readNum(buffer) : Function to read Numbers.
* readInt(buffer) : Function re read Integers.


## public.go
We will be creating a program public.go which has news data which is accessable to all. As its data is accessable to public, it has a very high risk that its content could be changed by some malicious actions. To prevent this we will checking news content before every serve. For verification we will be checking the hash values to be time efficient. If we found that the news that is being served to the user has been changed even slightly, the program will instantly restore it with the original content with the help of locally saved news content and then serve it to the user. 

### details about the program.
we have firstly stored the admin user id and password so that only he gets the authority to add new content and store the data locally. After that we have imported the locally saved csv data and stored it in LocalkeyValue hash map. . Functions used in public.go and their usage are:

* Put(key,val) : Take the key, value pair from the user and store it in the hash map (keyValue)
* save() : Function to save the content locally from hash map to csv. 
* Delete(key) : Function to delete the content(value and hash).
* Get(key) :  Function to retrieve news contnet and its Hash Value stored in a hash map.
* Put(key, value) : Function that takes key, value pair as an argument and store it in key,value and its hash256 in keyValue hash map.
* readString(buffer) : Function to read string.
* readNum(buffer) : Function to read Numbers.
* readInt(buffer) : Function re read Integers.
* check(usid,pass) : Function to check userid and password.
* Readcsv() : Function to readin CSV file(local data).
* put(key, value, hash) : Store key, value and hash, used in retrieving local data.
* Fingerprint(key, val) : Function to check the fingerprint and return a boolean result. That is true if mached and false if does not matched.
* update(key, val) : Update with new value for a key.


    
Also the admin can add the new news to the public data with the help of userid and password.

**UserId : arijit.Das**  
**Passwowd : UITBU.Arijit**
