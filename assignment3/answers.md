1. What are the boolean values in your language? (e.g., True and False, true and false, 1, and 0,  etc)

`true` and `false`

3. What types of conditional statements are available in your language? (if/else, if/then/else, if/elseif/else). Does your language allow for statements other than “if” (for example, Perl has  an “unless” statement, which does the opposite of “if”!)? What operators are used to act on booleans?

Conditionals:
    - `if`/`else`/`else if`
    - `:=`/`switch` and `case`
  
  Logical operators used to act on booleans:
    - `&&` and operator
    - `||` or operator
    - `!` not operator

  https://www.codecademy.com/learn/learn-go/modules/learn-go-conditionals/cheatsheet

3. How does your programming language deal with the “dangling else” problem?

Golang makes use of its garbage collecting feature and escape analysis to handle dangling else and other dangling pointers.

  https://web.archive.org/web/20200414222520/http://www.agardner.me/golang/garbage/collection/gc/escape/analysis/2015/10/18/go-escape-analysis.html

4. If your language supports switch or case statements, do you have to use “break” to get out of  them? Can you use “continue” to have all of them evaluated?
 
  The `break` statement is added automatically to a `switch` statement in Go and from my understanding `continue` statements are not supported or necessary.

  https://go.dev/tour/flowcontrol/9
  https://www.digitalocean.com/community/tutorials/how-to-write-switch-statements-in-go

5. Lastly, and perhaps most importantly: it is time to start thinking about what your programming project is going to actually be. In future sessions we will be breaking down your project into individual milestones for weekly check-ins. But for now, just write a short paragraph about what it is you are planning to code in your language of choice. Remember that it should be a substantial program - a game, app, calendar, web site, etc. (I'll try to dig up some examples to share before next class)

I am planning on designing a journal web app. This web app will have the following:
  - Allow users to create accounts
  - Different 'pages' for your journal that are set up like grids with 12 columns for each month and 31 rows for days
    - Each page will track a different aspect of daily life:
      - Excercise
      - Reading
      - Mood
      - Stress
      - And more
  - Also pages just to write something that made you happy/sad/something you did well and something you could've done better/etc.
  A user will create an account and have journal pages made available to them. Ideally, there will be default pages and then users can choose to make their own. As they fill in the days of their pages the grid will block out the day in a color (or possibly image) of their choice.
