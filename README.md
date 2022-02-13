SER Fyne Issue #2
=================

Possibly related issues: #1941, #165

### Describe the bug:

On Android, pressing backspace in an editable text entry widget deletes two characters.

### To Reproduce:

1. Clone this github repository
2. Package for Android
3. Push to an android device, install, and run
4. Select the text entry widget at the top of the screen
5. Enter some text (e.g., "abcdefg")
6. Press the keyboard's backspace (⌫) key

Two characters will be deleted. This happens from wherever in the string the cursor is, e.g. the end or the middle.

### Example code:

C.f. the repository supplied above. It is, essentially, a text entry widget with a list widget below. The text entry widget has an `OnSubmitted` lambda that prepends the text to the list.  The behavior happens without the `OnSubmitted` closure as well.

### Device (please complete the following information):

- **OS:** Android
- **Version:** Samsung One UI; Android 12. Built-in Samsung soft keyboard.
- **Go version:** 1.17.7
- **Fyne version:** v2.1.2
