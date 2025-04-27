# Special Note from carlRondoni

This repository is a fork of a kata project originally featuring multiple languages.
I tailored it by keeping only the languages relevant to our company’s stack — PHP, TypeScript — and added Go to expand our practice scope.
The objective was to streamline the repository for internal use and enhance the learning experience across our key technologies.

[Original repo](https://github.com/racingDeveloper/tell-dont-ask-kata)

# tell don't ask kata
A legacy refactor kata, focused on the violation of the [tell don't ask](https://martinfowler.com/bliki/TellDontAsk.html) principle and the [anemic domain model](https://martinfowler.com/bliki/AnemicDomainModel.html).

## Instructions
Here you find a simple order flow application. It's able to create orders, do some calculation (totals and taxes), and manage them (approval/reject and shipment).

The old development team did not find the time to build a proper domain model but instead preferred to use a procedural style, building this anemic domain model.
Fortunately, they did at least take the time to write unit tests for the code.

Your new CTO, after many bugs caused by this application, asked you to refactor this code to make it more maintainable and reliable.

## What to focus on
As the title of the kata says, of course, the tell don't ask principle.
You should be able to remove all the setters moving the behavior into the domain objects.

But don't stop there.

If you can remove some test cases because they don't make sense anymore (eg: you cannot compile the code to do the wrong thing) feel free to do it!

## Feedback for the creator of this forked version
Feedback is welcome!

How did you find the kata? Did you learn anything from it?

Please contact me on twitter [@racingDeveloper](https://twitter.com/racingDeveloper) or use the GitHub repo wiki!
