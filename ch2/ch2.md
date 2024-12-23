# Chapter 2

## Primitive Types and Declarations



### Built-in Types

```
"We’ll start by looking at primitive types and variables. While every programmer has
experience with these concepts, Go does some things differently, and there are subtle
differences between Go and other languages."
```



- Booleans, integers, floats and strings.
  
  Adopting Go's type system can be challenging for developers transitioning from other languages. Before diving deeper into specific types, let’s explore some foundational concepts that apply to all types.





### The Zero Value

In Go, variables that are declared but not explicitly assigned a value are automatically given a default zero value. This approach makes code more predictable and reduces potential bugs. Later on, we will also examine the zero value for each type.





### Literals

Literals represent numbers, characters, or strings directly in your code.

There are four common types of literals, along with a rare fifth type related to complex numbers (covered later).



#### Integer Literals

- Integer literals are sequences of digits that typically default to base ten. However, prefixes can specify other bases:
  
  - 0**b**: Binary (base 2);
  
  - 0**o**: Octal (base 8);
  
  - 0**x**: Hexadecimal (base 16);
  
  Avoid using a standalone `0` as an octal prefix (e.g., `075`), as it can be misleading and lead to confusion.
  
  In some programming languages, including Go, a number that starts with a `0` but has no other prefix (like `0x` or `0b`) is interpreted as an **octal literal**—a number in base 8. For example:
  
  - `075` in Go is not 75 in decimal (base 10). It is **61** in decimal because it is interpreted as an octal number (`7 * 8 + 5`).
  
  This behavior can be confusing because many developers from other languages (or those unfamiliar with octal literals) might mistakenly believe `075` is simply a decimal number. Using the explicit prefix `0o` for octal (introduced in Go 1.13) makes the intention clear:
  
  - **Good**: `0o75` clearly indicates octal.
  - **Confusing**: `075` may mislead readers into thinking it's a decimal number.
    
    

- **Note about octal numbers:**
  
  An **octal number** is a number expressed in **base 8**, meaning it uses only the digits `0` through `7`.
  
  In contrast, the decimal system (base 10) uses `0` through `9`, and the binary system (base 2) uses `0` and `1`.
  
  ##### How Octal Works
  
  Each digit in an octal number represents a power of 8, starting from the rightmost digit (similar to how decimal numbers represent powers of 10). For example:
  
  **Octal Number:** `173`  
  **Expanded Value in Decimal:**
  
  1×82+7×81+3×80=64+56+3=123 (in decimal)
  
  ##### Why Use Octal?
  
  Octal numbers are historically significant in computing and are often used in:
  
  1. **File permissions** in Unix/Linux systems (e.g., `chmod 755` uses octal).
  2. Representing groups of 3 binary digits (each octal digit corresponds to exactly 3 bits).
  
  ##### Example Conversion Between Systems:
  
  - **Binary (base 2):** `1101101`
  - **Group into 3s:** `1 101 101` → Add leading zero if needed: `001 101 101`
  - **Octal (base 8):** `1 5 5` → **Octal:** `155`
  
  This makes octal useful for simplifying binary data representation.
  
  In Go, you can use underscores to make long integer literals more readable. For example, `1_234` groups digits by thousands in base 10, or `0b1101_1010` separates bytes in binary. Underscores don’t affect the number's value but must not appear at the start, end, or next to each other. Use them wisely to improve clarity, like grouping by thousands in decimal or logical byte boundaries in binary, octal, or hexadecimal.



#### Floating Points Literals

- Represent numbers with fractional parts using a decimal point (e.g., `3.14`).

- You can also include an exponent using the letter `e`, followed by a positive or negative number (e.g., `6.03e23`).
  
  For example:
  
  - `6.03e23`: This means "6.03 multiplied by 10 raised to the power of 23." In this case, the decimal point shifts 23 places to the right, making the number much larger.
  
  - `6.03e-23`: This means "6.03 multiplied by 10 raised to the power of -23." Here, the decimal point shifts 23 places to the left, making the number much smaller (essentially a very small number).

- Alternatively, you can write floating point numbers in hexadecimal format by using the `0x` prefix and `p` to indicate the exponent (e.g., `0x1.23p4`).
  
  Let me break it down:
  
  - **`0x`**: This prefix indicates that the number is in **hexadecimal** (base 16) format.
  
  - **`1.23`**: This is the **fractional part** of the number in **hexadecimal**. Just like in the decimal system (e.g., `1.23` in base 10), `1.23` in base 16 means "1 plus 2/16 plus 3/256."
  
  - **`p`**: The letter `p` is used to indicate an **exponent** in this hexadecimal floating-point format, similar to how `e` is used in the decimal floating-point notation.
  
  - **`4`**: This is the exponent, specifying the power of 2 by which the number should be multiplied. In this case, `p4` means "multiply the number by 2^4" (or 16).
  
  So, `0x1.23p4` means:
  
  - First, take the fractional part `1.23` in hexadecimal (which is equivalent to `1 + 2/16 + 3/256` in decimal).
  - Then, multiply that result by `2^4` (or 16).
  
  This is a compact way to express floating point numbers in hexadecimal format, often used for more precise control in computations involving binary data.

- Just like integer literals, you can use underscores in floating point literals to improve readability (e.g., `1_234.56`).





#### Rune Literals

- Used to represent characters and must be enclosed in single quotes. 

- You can represent these characters in various ways, such as using Unicode values, octal, or hexadecimal numbers. 

- Go also supports escape sequences for special characters like newlines and quotes. 

- Unlike many languages, Go distinguishes between single and double quotes, reserving single quotes for characters (runes) and double quotes for strings.
  
  ##### Ways to represent rune literals:
  
  1. **Single Unicode characters**: A rune literal can be a single Unicode character, for example, `'a'`.
  
  2. **8-bit octal numbers**: A rune can also be written using an **8-bit octal** representation, such as `'\141'`. This represents the ASCII value of `'a'` in octal.
  
  3. **8-bit hexadecimal numbers**: Similarly, a rune can be expressed as an **8-bit hexadecimal** number, for example, `'\x61'`. This is the hexadecimal equivalent of `'a'`.
  
  4. **16-bit hexadecimal numbers**: For characters beyond the ASCII range, you can use **16-bit hexadecimal** Unicode values. For example, `'\u0061'` represents `'a'` in 16-bit Unicode.
  
  5. **32-bit Unicode numbers**: You can also use **32-bit Unicode** values to represent characters, such as `'\U00000061'` for `'a'`.
  
  ##### Escape sequences for rune literals:
  
  Go provides a number of **escape sequences** for representing special characters:
  
  - `\n`: Newline character.
  - `\t`: Tab character.
  - `\'`: Single quote (to escape single quotes).
  - `\"`: Double quote (to escape double quotes).
  
  For example:
  
  Without escaping: `"He said, "Hello!"."` — This will cause an error because the compiler thinks the string ends at the first `"`.
  
  With escaping: `"He said, \"Hello!\"."` — The backslash `\` escapes the double quotes, so the string can include quotes within it.
  
  - `\\`: Backslash (to escape backslashes).



#### String Literals

##### Interpreted Sting Literals

- These are created using **double quotes** (e.g., `"Hello"`).

- You can include special characters, such as escape sequences like `\n` for a newline or `\"` for a double quote inside the string.

- However, there are some things you **cannot** do directly inside an interpreted string:
  
  - You can't include an **unescaped backslash** (like `\` without something following it).
  - You can't have **unescaped newlines** (the string can't break onto the next line unless you use `\n`).
  - You can't include **unescaped double quotes**.

- Example:
  
  - To have "Greetings" on one line and `"Salutations"` (in quotes) on the next line, you would write:
    
    ```
    "Greetings and\n\"Salutations\""
    ```
  
  - This gives:
    
    ```
    Greetings and 
    "Salutations"
    ```

##### Raw String Literals

- These use **backquotes** (`` ` ``) instead of double quotes.

- You can write strings exactly as they appear, without worrying about escape sequences for newlines, backslashes, or quotes (except for the backquote itself).

- Raw strings are helpful for multi-line strings or strings that include special characters like backslashes and quotes.

- Example:
  
  - If you want to write a multi-line string like:
    
    ```
    Greetings and 
    "Salutations"
    ```
  
  - You can use a raw string literal like this:
    
    ```
    `Greetings and
    "Salutations"`
    ```



> Go allows integer literals to be used in floating point expressions or assigned to floating point variables because **literals are untyped**. 
> 
> 
> 
> This means they can work with any compatible variable. However, **you can’t assign a literal string to a numeric variable or a float literal to an int, as these will cause compiler errors. While literals are untyped to provide flexibility, Go enforces type compatibility**.
> 
> 
> 
> For example, assigning a value that overflows the target variable’s size, like 1000 to a byte, will result in a compile-time error. **When the type isn't explicitly declared, Go defaults to a type for the literal**, we'll explain it better later on variable assignment section.
































