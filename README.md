# Calendar Application – Requirements Specification

## 1. Overview
The task is to develop a **simple graphical calendar application**.  
The application displays the days of a selected month and year, allows navigation through months and years, highlights Sundays and holidays, and reads holiday data from a file at startup.

No additional functionality is required beyond displaying and closing the calendar.

---

## 2. Functional Requirements

### 2.1 Calendar Display
- The calendar displays **all days of the currently selected month and year**
- Days are displayed in a **grid layout**, with **7 days per row**
- The start of the week must be consistent (Monday–Sunday)
- **Sundays must be visually highlighted**
- When starting, it should display the current month

### 2.2 Month and Year Selection
- The **month** is selected using a **ComboBox (drop-down list)**
- The **year** is entered manually in a **text input field**
- When the month or year changes, the calendar **updates immediately**

### 2.3 Date Jump Functionality
- A separate input field allows entering an **arbitrary date**
- Entering a valid date causes the calendar to **jump to the corresponding month and year**
- The supported date format will be `DD.MM.YYYY`
- Invalid date input should be handled gracefully (e.g. ignored or reported)

### 2.4 Holidays
- On application startup, the program reads an **ASCII text file** containing holidays
- Each entry in the file includes:
  - a **date**
  - a flag indicating whether the holiday is **recurring** (e.g. yearly)
- The date and recurring flag are separated by an **arbitrary delimiter** 
- The file does **not** need to contain all holidays; a few examples are sufficient
- **Holidays must be visually distinguished** from normal days (different from Sunday highlighting)
- If a holiday is on Sunday, it should also differ from Sundays or other holidays
- Invalid entries (rows) in this file should be skipped

### 2.5 Restrictions
- No additional functionality is required except **closing the application** (window exit button will do)
- The calendar relies on Go’s standard time library, which only supports dates between years 1 and 9999, the year range is therefore intentionally limited to ensure correctness

