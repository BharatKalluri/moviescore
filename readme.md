# MovieScore

## Show me what it does!
[![asciicast](https://asciinema.org/a/1ms6yrsd8tryotybmpnm5vsaj.png)](https://asciinema.org/a/1ms6yrsd8tryotybmpnm5vsaj)

## What Does It Do?

A Cli utility written in go language to get movie ratings , Trailer Links and reviews from the internet. Right into your Terminal!

It gets A whole lot of information about the Movie such as
- Director
- Cast
- Year
- Release Date
- Rated
- Genre
- Poster
- Metascore ratings
- Awards
- Plot
- Movie Trailer
- IMDB rating
- Rotten Tomatoes Rating

and with another flag, it get reviews from Rotten tomatoes!

## How to Use it?

Its Super Simple!

For all the movie information
```bash
moviescore "The Dark Knight"
```
will give you the following
```
------------------------------------------------------
  __  __            _         _____                    
 |  \/  |          (_)       / ____|                   
 | \  / | _____   ___  ___  | (___   ___ ___  _ __ ___ 
 | |\/| |/ _ \ \ / / |/ _ \  \___ \ / __/ _ \| '__/ _ \
 | |  | | (_) \ V /| |  __/  ____) | (_| (_) | | |  __/
 |_|  |_|\___/ \_/ |_|\___| |_____/ \___\___/|_|  \___|
------------------------------------------------------

 Movie Name: The Dark Knight
 Director: Christopher Nolan
 Cast: Christian Bale, Heath Ledger, Aaron Eckhart, Michael Caine
 Year: 2008
 Released: 18 Jul 2008
 Rated: PG-13
 Genre: Action, Crime, Drama
 Poster: https://images-na.ssl-images-amazon.com/images/M/MV5BMTMxNTMwODM0NF5BMl5BanBnXkFtZTcwODAyMTk2Mw@@._V1_SX300.jpg
 Metascore Rated: 82
 Awards: Won 2 Oscars. Another 146 wins & 142 nominations.
 Plot: Set within a year after the events of Batman Begins, Batman, Lieutenant James Gordon, and new district attorney Harvey Dent succ
essfully begin to round up the criminals that plague Gotham City until a mysterious and sadistic criminal mastermind known only as the 
Joker appears in Gotham, creating a new wave of chaos. Batman's struggle against the Joker becomes deeply personal, forcing him to "con
front everything he believes" and improve his technology to stop him. A love triangle develops between Bruce Wayne, Dent and Rachel Daw
es.
 Movie Trailer: https://www.youtube.com/embed/YoHD9XEInc0
 Ratings from IMDB and Rotten Tomatoes---
 IMDB Rating: 9.0
 Rotten Tomatoes Rating: 94% (Certified Fresh!)

```

For reviews
```bash
moviescore -r "The Dark Knight"
```
will give you
```
------------------------------------------------------
  __  __            _         _____                    
 |  \/  |          (_)       / ____|                   
 | \  / | _____   ___  ___  | (___   ___ ___  _ __ ___ 
 | |\/| |/ _ \ \ / / |/ _ \  \___ \ / __/ _ \| '__/ _ \
 | |  | | (_) \ V /| |  __/  ____) | (_| (_) | | |  __/
 |_|  |_|\___/ \_/ |_|\___| |_____/ \___\___/|_|  \___|
------------------------------------------------------

 Reviews from rt 
 ------------------- 
The Dark Knight is probably the smartest and most stylish action movie since the "The Matrix." It thinks and philosophizes. The subject
 it thinks and philosophizes about, perhaps not surprisingly, considering the times, is the Iraq war.
 ------------------- 
An exceptionally smart, brooding picture with some terrific performances.
 ------------------- 
and many more(shortened for readme)........
```
For some A year argument is required. For example searching for arrival, rotten tomatoes rating will not be shown and reviews will not work. In that case
```bash
moviescore -y 2016 "arrival"
```
will work!

## How to install

If you have go language on your system, then
```bash
go get github.com/bharatkalluri/moviescore
cd $GOPATH/src/bharatkalluri/moviescore
go install
```
and you are good to go!

If you don't have go installed on your system , Then download the moviescore file from dist folder and place it in a convenient folder, then it can be used from there..
```bash
./moviescore inception
```

### Have any ideas?
Feel free to contact me at kalluribharat@gmail.com
