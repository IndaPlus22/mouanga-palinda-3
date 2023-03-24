If written answers are required, you can add them to this file. Just copy the
relevant questions from the root of the repo, preferably in
[Markdown](https://guides.github.com/features/mastering-markdown/) format :)


Take a look at the program [matching.go](src/matching.go). Explain what happens and why it happens if you make the following changes. Try first to reason about it, and then test your hypothesis by changing and running the program.

  * What happens if you remove the `go-command` from the `Seek` call in the `main` function? 
  
  **Hyp**: the program becomes deterministic and A&B and C&D match but E doesn't

  **Ans**: what i said


  * What happens if you switch the declaration `wg := new(sync.WaitGroup`) to `var wg sync.WaitGroup` and the parameter `wg *sync.WaitGroup` to `wg sync.WaitGroup`?

  **Hyp**: either nothing, or a compile error because the pointers aren't pointy enough

  **Ans**: we're passing the lock by value now instead of by reference (in Seek), so we're modfying the wrong lock, leading to a deadlock

  * What happens if you remove the buffer on the channel match?

  **Hyp**: the last call to Seek doesn't receive anything from match, and doesn't call wg.Done(), and the program never ends (or crashes because all routines are asleep)

  **Ans**: deadlock, because of the unbuffered channel in which the last message is never received

  * What happens if you remove the default-case from the case-statement in the `main` function?

  **Hyp**: program doesn't end since it's waiting for something from the match channel
  
  **Ans**: there was one name left in the channel and we selected it, after that the main fn ended so the program ended.

Hint: Think about the order of the instructions and what happens with arrays of different lengths.

