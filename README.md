# Quiz Master App
Is a simple quiz master app using [Cobra CLI](https://github.com/spf13/cobra) with [sqlite db](https://www.sqlite.org/index.html) for embedded local database storage for all the questions.

# How to setup the quiz_master app?
1. Clone this repo to your local machine and compile it accordingly based on your OS platform.
2. If you have Make software installed, you can easily run it with this command "make build" or "make all" to setup, build, test
3. If all is well, you can run the quiz_master app in your local machine.

# App Features
1. Main app interface
![image](https://user-images.githubusercontent.com/58651329/149918286-038a8fa5-6f71-4608-a717-9d92e1d0ae21.png)

2. Help command, for e.g: ```quiz_master create_question --help```
![image](https://user-images.githubusercontent.com/58651329/149918504-05c1c0ae-0fbf-445f-ab24-5ba9fd697791.png)

3. Command to create a new question, ```quiz_master create_question 1 "How many letters are there in the English alphabet?" 26 ```
![image](https://user-images.githubusercontent.com/58651329/149918050-9fb9e6e0-db90-4e12-8634-3c0105f27cc4.png)

4. Command to delete a question, ```quiz_master delete_question 1```, there will be a prompt first to confirm user's action.
![image](https://user-images.githubusercontent.com/58651329/149919237-ca1d062c-fef9-476e-b709-62095bb5325f.png)
![image](https://user-images.githubusercontent.com/58651329/149918948-1631cc04-3153-4478-9744-d318da789267.png)

5. Command to update the question, ```quiz_master update_question 1 "Test update" 25```, if the question number is not found, throw and error instead.
![image](https://user-images.githubusercontent.com/58651329/149919599-d0fce4dc-4756-42ed-9fe3-fba3946f7190.png)

6. Command to display all questions information, ```quiz_master questions```
![image](https://user-images.githubusercontent.com/58651329/149920304-108fe100-12b2-47c1-8974-d09399d840f8.png)

7. Command to display a single question with inputs from the user's for their answer, ```quiz_master question 1```
![image](https://user-images.githubusercontent.com/58651329/149920687-949bec13-c0ad-4bd6-9e65-7f09b87cd505.png)
![image](https://user-images.githubusercontent.com/58651329/149920913-659eb7ea-012c-452c-85a0-1a3cf4112f1a.png)
