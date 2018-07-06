# sgrda

Search Go Repository Dependency Application 

Return repositories from variadic dependencies.


Gathering: Webcrawlers
    From: Github, bitbucket,
    How often?
    With? cully, go-query, 
    ToCrawl? What to crawl and how is it triggered
    Add to cassandra Db
    api repo imports from  http://go-search.org/ 


Client Interaction:
    Search for repos with dependencies inputed. 
    Query against cassandra. 
    Return list of repos with all dependencies.