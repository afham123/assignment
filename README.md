# assignment

## Hashing
Hashing is the process of scrambling a piece of information or data beyond recognition. We use hash function to convert input into hash Digest. These functions are irreversible by design. The value generated after passing the plane text information through the hash function is called the hash value digest or just the hash of the original data.  Hashes are made to be irreversible and no discription key can convert a digest back to its original plane text value. There are few hasing algorithms that still stands undecodable and are used for password storage and integrity verification, etc. 

In this project we will be using hashing to insure our data intigrity. We will initially pass our original data to hashing function which grenerate a hash value and it will be stored in our data base in the severs. If any changes or modification will be their in the data we will recalculate the hash value of modified data and if we does matches our hash value which as been stored we will restore back to its initiall form and in this way we can insure that nobody could modify or delete our data. 

![Image](https://user-images.githubusercontent.com/49563140/136780339-4bc8e5cb-7bbd-4be9-85e4-ed3b8a0d7f67.png)


### Hashing algorithms(SHA 256)

We will be using SHA 256 algorithm for  implement an efficient in-memory Verifiable. The algoritm will prduce a digest is of length 256 bits. The digest produce will be irreversible that is no method or function will be able to retrieve the original text even when the digest is passed through the hash same hash function. In this way we can create a digest of all original news content and store it in our database. If any changes will be made in the news paper its digest is mached with the original news digest and if we found the digest to be insimilar. We could be able to know that the news contnet have been chnaged. In that case we can resore it to original content.
![Image1](https://user-images.githubusercontent.com/49563140/136897940-b9b45c00-bc9f-47bf-9815-87f9442d1406.png)
