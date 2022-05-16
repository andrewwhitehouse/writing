[_My brain has been jumping around this morning. I started off thinking about what I can do to support my daughter in preparing for her summer internship. And then I thought I'm second-guessing this and I should do the simplest thing and have sessions with her and talk through what she wants to learn. And I have this book I want to write, which could be helpful to her, and to others. And leverage the work I've already done ..._]

**The Elements of Programming**

The Structure and Interpretation of Computer Programmes is one of the first texts I read, and the one book I kept from my university study. It was written in Scheme, which is a form of LISP.

I've decided to keep one eye on this book, cherry-picking, and updating the examples to align with modern languages currently in use.



From [SICP](https://mitpress.mit.edu/sites/default/files/sicp/full-text/book/book.html):

> A powerful programming language is more than just a means for instructing a computer to perform tasks. The language also serves as a framework within which we organise our idas about processes. Thus, when we describe a language, we should pay particular attention to the means that the language provides for combining simple ideas to form more complex ideas. 
>
>Every powerful language has three mechanisms for accomplishing this:
>
>- primitive expressions, which represent the simplest entities with which the language is concerned
>
>- means of combination, by which compound expressions are built from simpler ones, and
>
>- means of abstraction, by which compound objects can be named and manipulating as units.
>
>In programming we deal with two kinds of objects: procedures and data. [...] Thus, any powerful programming language should be able to describe primitive data and primitive procedures and should have methods for combining and abstracting procedures and data.
>
>In this chapter we will deal only with simple numerical data, so that we can focus on the rules for building procedures.

I've spent more than 30 years working as a software developer. I've read a lot of books, which tell you lots of stuff and get you to write code. I rarely feel that these get to the essence of what it is we're doing ... why a particular language was designed the way it was and what are the pros and cons of doing it that way. So rather than just following a series of steps with some overly simplified statement "comments are useful", I want to avoid wasting people's time _and_ not mislead them.

As David Foster Wallace reminds us in [This Is Water](https://fs.blog/david-foster-wallace-this-is-water/) it's helpful to acknowledge what ocean you're swimming in.

[This may belong in an introduction ...]

One of the things I remember my Dad making for me  was a ply wood box on which were mounted lots of different types of switches (push and release, sliding, chrome toggle switches) and lamps like you might find in an aircraft cockpit. After playing with it for a while I opened it up to see how it works. 

This was the analogue world, which obeys the laws of physics. If the battery started to drain the bulb glowed less brightly. But it was "on" or "off". 

In the "digital" world we map those analogue voltages to binary 1s (on) and 0s (off): on is represented as 5 volts usually, and off is zero. You'll see that your phone charger delivers 5 volts.

These electrical signals were initially manipulated by valves and then transistors. To increase the speed and capability of the processors, chip manufacturers packed more and more transistors on to the silicon wafer. Until the laws of physics meant that they couldn't do that reliably on a single processor. So now we have multiple processors.

Our primary school headmaster suggested that a boarding school in our county might be a good option. I continued my interest in electronics, buying books and the magazine Everyday Electronics, and figuring out how to organise components on veroboard to match the circuit diagram.

At boarding school there was a minicomputer which ran a version of Unix. The programming language was C, and so I bought a copy of the (pre-ANSI) "C Programming Language" and worked through the examples.

This book is useful because it gives you experience coding closer to the hardware. 

As coders we operate some way up the stack, and we don’t usually need to think about the details of what the hardware is doing to make our programme run.

We can write code that uses the fewest possible machine instructions to achieve the result, if performance is really critical. But that requires an investment in understanding how our chosen language maps to machine instructions. And that may not be the most valuable use of our developers’ time compared to shipping features for our customers who pay the bills.

Our programme runs on a computer that has a finite amount of memory to store its working calculations. If we exceed that memory our programme will stop working properly. So something, or someone, needs to clean up the unused memory. Garbage collectors do this in many language runtimes. Or you can avoid the overhead of the cleanup and write code that is careful to allocate only what is needed, and clean it up.

The instruction, or word size, affects the data types that are available to you as a coder.

We used to have “32-bit” computers which referred to the instruction size. The maximum number that you can natively represent in the hardware was 2^32, which is 4,294,967,296. So 32 bit computers could not have more than 4GB (4 gigabytes … giga is 10^9). Now that we have 64-bit machines, the hardware can support much larger amounts of memory.

(Note: calculate 2^64 in the next REPL session. And note that Python doesn’t limit integers to this size.)

In the course of your coding you will create variables to represent the data concepts you are modelling. A statically typed language like Rust requires you to specify up-front what the type is, including how big you need it to be. A dynamically typed language infers the type, which means you as the coder don’t need to think about this too much.

It's helpful to remembers that we have data abstractions built on top of the ones and zeroes recognised at the hardware level. Because sometimes those abstractions are "leaky" and we need to deal with that (example: how to represent floating point numbers).

To compare and contrast language styles we'll be covering examples in Python, Clojure and Rust.
