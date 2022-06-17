**Language Evolution**

[_I've been watching The Staircase. I noted how the film makers collected a load of footage and then debated at length what to leave in and what to leave out ... whether it make him seem more or less guilty. This writing process feels like that. There are a couple of things I want to address in this post and I don't feel like I've nailed it yet: how programming languages and their adoption evolves ... from assembly to C to high level languages; and also abstractions ... while the focus is on a specific language it's important to teach transferable principles, like abstactions, which can still be applied when the reader learns a different programming language._]

I wrote my first substantial programme when I was 15.

With my Dad having passed away when I was four, and coming from a single parent family, my Mum decided it would be better for us to go to a boarding school.

I had an interest in electrical and electronic things early on. I used to buy the magazines Practical Electronics and Everyday Electronics and make some of the sample circuits. 

And the school had early computers, including that accepted their programmes via punched cards. The hole (or absence of a hole) represents a binary 1 or 0 and when there are eight of them you can represent up to 2^8 or 256 values. The values can represent a type of instruction (like "increment") and a memory location, or register, to operate on.

Assembly code runs quickly; you are coding very close to the hardware and so there is no extra work to be done. But coding at this low level requires you to think at this low level, and write (and maintain) lots of code. Conditional logic is achieved by "jumping" to different parts of the programme, and you may need to specify those locations directly. So when you add or remove code you need to adjust the "distances".

My first language was C, which is also a "low-level" language. C introduces the data types, conditional logic and [add other things] which allow the developer to express the programme at a higher than assembly language. But one thing that C requires you to do is manage your own memory. As your programme runs you create temporary data and if your programme is running for long enough you need to make sure you clean up when you've finished using that memory. Otherwise you will eventally run out of available memory. It's also up to you to ensure that memory is allocated before you use it. If you try to access memory that hasn't been allocated correctly your programme will fail.

Programming languages come in and out of fashion. Some are long-lived, particularly if they have a strong community around them, and an effective way of deciding how to evolve. 

The challenges of memory management in C have given rise to other languages. Java has similar syntax to C, and (mostly) takes care of the memory managament for you. It is still widely used in business applications. Some have seen ways that it can be improved, by being more concise, and not requiring an object-oriented style of programming. And while the automatic garbage collection takes away the overhead from the programmer, it can introduce latency (delay) in executing your programme.

Go is a modern language that was designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. Ken Thompson was also one of the designers of the C language.

[Add more justification for Go here.]

Go is the lanugage we're going to use in this introduction to coding. 

We're going to focus on the code ... more showing and less telling. 

One of my favourite places to visit when I was young (and now) is the science museum. I love the interactive element of the science museum where you pressed a button and could observe the result. Sound doesn't travel through a vacuum, and you can see it for yourself.

For those of us who prefer to learn kinaesthetically, this book will give you plenty of examples as a way to building mental models about coding, and show the abstractions that we use to construct software systems.


