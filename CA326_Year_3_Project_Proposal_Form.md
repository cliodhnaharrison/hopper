# CA326 Year 3 Project Proposal Form
<b>SECTION A</b>
Project Title: Hopper  
Student 1 Name: Maciej Swierad     ID Number: 18455726  
Student 2 Name: Clíodhna Harrison     ID Number: 15440568  
Staff Member Consulted: Stephen Blott  

## Project Description (1-2 pages):



<b>Description</b> - Minimum 250 word description of the proposed project.
- Our  idea is to provide students with discardable containers in which they have full access to play around as root without any of the drawbacks playing around as root would have on an actual system or even a VM. The containers would be discarded as soon as the student disconnects.
- For a module like CA282 DevOps students could do what is required of them for a lab in the container and then the logs/history could be evaluated by a script elsewhere or possibly they could place their work in a volume mounted to the container where a script could evaluate their work. 
- Regardless of how the student’s work is evaluated their history/logs will be sent to them afterwards for their records.
- Students would connect to the service by logging into a web terminal in their browser, this means that any student regardless of their operating system can use the containers. Which would help with the general accessibility of the module. 
- Students are unable to upload to a web terminal so the possibility of plagiarism is reduced. Beyond that an anti-cheat script could check what commands a student runs and in what order to give a plagiarism score similar to TurnItIn does for essays.
- Back end will be written mostly in Go, it’s purpose will be to spin up containers for each student on demand. The server side code will handle allocation of resources similar to orchestration software. Simultaneously the back end will be responsible for grading the student’s work.


<b>Division of Work</b> - Outlines how the work is envisaged to be split equally among the team members.
We plan to work together on most parts of the project, we can’t see any easy split in the project and want both of us to learn from all parts of the project and be familiar with all parts. We will be employing tactics such as  pair programming which we have become quite familiar with over the previous college years. We will be working together over the likes of voice channels such as Discord wherever possible.


<b>Programming language(s)</b> - List the proposed language(s) to be used
Go, Bash

<b>Programming tool(s)</b> - List tools (compiler, database, web server, etc.) to be used
Makefiles, Docker, web terminals

<b>Learning Challenges</b> - List the main new things (technologies, languages, tools, etc) that you will have to learn
Neither student has experience with Go. Both students have minimal experience with Docker. Maciej has experience with Bash and minimal experience with Makefiles. Clíodhna has experience with Makefiles and minimal experience with Bash.

<b>Hardware / software platform</b> - State the hardware and software platform for development, eg. PC, Linux, etc.
Linux, WebApp 

<b>Special hardware / software requirements</b> - Describe any special requirements.
N/A
