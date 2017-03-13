# MovieScore

[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php)
[![Open Source Love](https://badges.frapsoft.com/os/v3/open-source.png?v=103)](https://github.com/ellerbrock/open-source-badges/)
[![Go Reportcard](https://goreportcard.com/badge/github.com/bharatkalluri/moviescore)](https://goreportcard.com/report/github.com/bharatkalluri/moviescore/)
[![Bash Shell](https://badges.frapsoft.com/bash/v1/bash.png?v=103)](https://github.com/ellerbrock/open-source-badges/)

## Show me what it does!
[![asciicast](https://asciinema.org/a/106937.png)](https://asciinema.org/a/106937)

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

with the -r flag
- Rotten tomatoes reviews

with the -pg flag
- Parental guide from IMDB

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

For Parental Guide
```bash
moviescore -pg inception
```
will result
```bash
------------------------------------------------------
  __  __            _         _____                    
 |  \/  |          (_)       / ____|                   
 | \  / | _____   ___  ___  | (___   ___ ___  _ __ ___
 | |\/| |/ _ \ \ / / |/ _ \  \___ \ / __/ _ \| '__/ _ \
 | |  | | (_) \ V /| |  __/  ____) | (_| (_) | | |  __/
 |_|  |_|\___/ \_/ |_|\___| |_____/ \___\___/|_|  \___|
------------------------------------------------------

 Parental Guide:
Reference Url:  http://www.imdb.com/title/tt1375666/parentalguide
Sex and Nudity 
  No sex or nudity is shown.Cobb and Saito first meet in a "love nest" where Saito frequently meets his lover - a relationship he's
  managed to keep secret in his waking life.We see people kiss, and one character uses subterfuge to steal a kiss.A couple of women
  wear low-cut tops.
Violence And Gore
  The only blood shown in this film is when a man is shot in the chest about halfway through the movie, where we see red seep through a
  man's shirt as he coughs up flecks of bloodViolence in Inception is tricky to tally. At times we see real men get hurt or killed. But
  much of the violence is perpetuated in dream worlds, where the people we see are not real, but manifestations of the subject's
  subconscious. As a result, the "real" body count is surprisingly low (at least for a film that wields this much intensity), while
  metaphysical fatalities run off the chart.Merging both categories, people are punched, kicked, choked, shot (scores of times),
  stabbed, hit by cars (several times), blown up, attacked by rampaging mobs, almost buried by avalanches and nearly drowned. Somebody
   gets shot in the foot just to illustrate that, while dying in a dream state is difficult, pain is all too easy to come
  by.The visceral feel of the violence is about what you'd expect for a PG-13 movie and, frankly, maybe a step back from a
  prime-time actioner on television. The mayhem is practically bloodless, and it's perpetrated with a certain, almost chilly,
  remove.When someone wants to exit a dream, they simply "kill" themselves or have someone do it for them. Cobb, for instance,
  shoots one of his compadres in the head to wake him up. (We see a bloodless hole in his forehead briefly.)Because the
  sensation of falling can jar someone awake, folks routinely engineer the end of their dreams by plummeting off bridges,
  jumping out of tall buildings or cutting loose elevator cables. One character throws another off a cliff.A scene is shown
  several times where a man and woman are lying on a train track with the intention of committing suicide.
 Profanity
  8 uses of hell, 4 uses of god damn it , 2 uses of Jesus Christ, 2 uses of asshole , 2 uses of damn, 2 uses of bloody,2 uses of for
   god's sake,2 uses of Jesus, 1 use of bastards, 1 use of god damn.
Alcohol/Drugs/Smoking
  People drink wine and beer. Intravenous drug sedatives and other mysterious concoctions are required to put people into these dreamlike
  states.A character has lost his ability to dream normally, and so he repeatedly hooks himself up to delve into his own haunting dream
  world.
 Frightening/Intense Scenes
  This film has an intense atmosphere throughout. The second half of the film is especially intense.In a dream, people hold a woman
  while another woman walks up and stabs her. She wakes up before you can see any harm done. This scene is unexpected and can frighten
  some.A woman shoots a man's leg to torture him.A train unexpectedly crashes into a car with people in it.A street starts
  exploding.(in a
  dream)A woman jumps off a building, believing it to be a dream and thinking she will wake up. (The concept of her believing the
  world to be a dream, and that she has to kill herself, can be quite frightening)The action is intense at times. People die. (see
  violence/gore)The concept of the dream-levels in the movie, and the fact that you might not be sure if the characters are dreaming
  or not, might frighten some viewers.The suicide of a character because he/she thought he/she were in a dream may be disturbing to
  some viewers.BBFC - (12A) Moderate violence.
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
cd $GOPATH/src/github.com/bharatkalluri/moviescore
go install
```
and you are good to go!

If you don't have go installed on your system , Then download the moviescore file from dist folder and place it in a convenient folder, then it can be used from there..
```bash
./moviescore inception
```

### Have any ideas?
Feel free to contact me at kalluribharat@gmail.com

###### Note
All the information belongs to the respective owners. The information is just being scraped from the publicly available web pages.
