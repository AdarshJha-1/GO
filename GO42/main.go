package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// f, err := os.Open("letters.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// n, err := countAlphabets(f)

	r := strings.NewReader("Omae Wa Mou Shindeiru") //Y ou Are Already Dead

	n, err := countAlphabets(r)
	if err != nil {
		panic(err)
	}

	fmt.Println("Letters:", n)

	newFile, err := os.Create("writing.txt")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	content := `I dreamt to be the very best
Among the strong and brave
But masters of the Pokémon
Prefer to feel like slaves
I left my home, I tried to pass
Right through the tall low poly grass
As the strangest feeling in my pants
Came like a tidal wave
A wild Lopunny appears
She's readily spreading her ears
You see an Incineroar flexing, oh, dear
This ecosystem is weird
Heed their call to catch 'em all
Master, stop rubbing your blue poké balls
Teach us new moves and carry us home
We always do as we're told
The furrýmonication's growing bolder
They say it's in the eye of the beholder
I say it's just effective
Yeah, it's super effective
More than a couple extra polygons
My Fennekin evolved into a furry
An animal or not? The line gets blurry
I guess this is effective
Yeah, it's super effective
So grab your ultra balls
Gotta smash 'em all!
Pretend you like this game a lot
Pretend you play it for the plot
The only rule that you've been taught
Must be the 34th
You see Machokes, they roam the wild
It looks disturbin', but you don't mind
'Cause in a fanfic that you write
It's getting so much worse
Pondering hearts, hardening parts
Your partners completed the training
Your battling art's defying the charts
It's so much more entertaining
Gym leader appears, the audience cheers
Shattering battle arena
Well, she has no clue what you gonna do to her naive Primarina
Enemy Pokémon fainted
(Ooh!) This pleasure, oh, it's so tainted
Look at her trainer, she's devastated
Take this, you haters of furry baiting
A silence fell over the crowd
(Ooh!) Only some kid is crying aloud
But you won the match and you got the badge
It's over the edge, but your mommy is proud
Why not just cast a line and see who's baited?
Reshaping the designs full-on R-rated
Well, that could be effective
Yeah, it's super effective
Why polish game mechanics or a plot
And at this point, who cares what I've been drinking?
I find the Eeveelutions pretty kinky
It really is effective
Yeah, it's super effective
So don't forget your goal
Gotta smash 'em all!
Lucario used "Choking"
And you don't wanna stop him
A route to any fetish
A harem in your pocket
Scorbunny likes it faster
Machoke is gonna rock it
Who wouldn't wanna master
A harem in your pocket (gotta grab 'em all)
If you are into felines
You better ask Team Rocket
A route to any fetish
A harem in your pocket (gotta yiff 'em all)
Scorbunny likes it faster
Machoke is gonna rock it
Who wouldn't wanna master
A harem in your pocket (gotta breed 'em all)
Yeah, you should be a master
A harem in a pocket
A furry waifu's master
A harem in your pocket (gotta smash 'em all)
Yeah, you should be a master
A harem in a pocket
A furry waifu's master (ooh, yeah)
A harem in your pocket (gotta smash 'em all)`
	n, err = writeString(content, newFile)
}

func countAlphabets(r io.Reader) (int, error) {
	count := 0
	buffer := make([]byte, 1024)

	for {
		n, err := r.Read(buffer)
		for _, l := range buffer[:n] {
			if (l >= 'A' && l <= 'Z') || (l >= 'a' && l <= 'z') {
				count++
			}
		}

		if err == io.EOF {
			return count, nil
		}
		if err != nil {
			return 0, err
		}
	}
}

func writeString(s string, w io.Writer) (int, error) {
	n, err := w.Write([]byte(s))
	if err != nil {
		return 0, fmt.Errorf("error occurred while writing: %w", err)
	}
	return n, nil
}
