#     ___   __   __    __     __   _  _   __   _  _    ____  _  _  ____  ____  __   __ _       ____   __  ____  ____ 
#    / __) / _\ (  )  (  )   /  \ / )( \ / _\ ( \/ )  / ___)/ )( \(_  _)(_  _)/  \ (  ( \ _   (___ \ /  \(___ \( __ \
#    ( (__ /    \/ (_/\/ (_/\(  O )\ /\ //    \ )  /   \___ \) \/ (  )(    )( (  O )/    /( )   / __/(  0 )/ __/ (__ (
#    \___)\_/\_/\____/\____/ \__/ (_/\_)\_/\_/(__/    (____/\____/ (__)  (__) \__/ \_)__)(/   (____) \__/(____)(____/

#     SOFTWARE WARRANTY LICENSE NOTICE
#     (c) 2023 Calloway Sutton. All rights reserved.

#     This software is provided "as is" and without any express or
#     implied warranties, including, but not limited to, the implied
#     warranties of merchantability and fitness for a particular
#     purpose. In no event shall the authors or copyright holders
#     be liable for any claim, damages, or other liability, whether
#     in an action of contract, tort, or otherwise, arising from,
#     out of, or in connection with the software or the use or other
#     dealings in the software.

#     For more information, please contact: me@callowaysutton.com

from reportlab.pdfgen import canvas
from reportlab.lib.pagesizes import letter
from reportlab.lib.utils import simpleSplit
import random
import tqdm
from lorem_text import lorem
from multiprocessing import Pool
from pathlib import Path

num_pdfs = 100
process_count = 8  # Number of processes to use

def draw_text_with_word_wrap(c, x, y, text, width):
    lines = simpleSplit(text, c._fontname, c._fontsize, width)
    c.setFont("Helvetica", 12)  # Adjust the font name and size as needed
    c.drawString(x, y, lines[0])
    c.setFont("Helvetica", 10)  # Adjust the font name and size as needed
    for line in lines[1:]:
        y -= 14  # Adjust the line spacing as needed
        c.drawString(x, y, line)

def generate_pdf(i):
    name = lorem.words(8).replace(" ", "_")
    filename = f"pdfs/{name}_{i}.pdf"
    doc = canvas.Canvas(filename, pagesize=letter)

    num_sentences = random.randint(1000, 2000)

    for _ in range(num_sentences):
        sentence = lorem.paragraphs(5)
        draw_text_with_word_wrap(doc, 50, 700, sentence, 500)  # Adjust the position and width as needed
        doc.showPage()

    doc.save()

if __name__ == "__main__":
    # Create the output directory if it doesn't exist
    Path("pdfs").mkdir(exist_ok=True)

    # Create a pool of processes
    with Pool(process_count) as pool:
        # Use the pool to generate PDFs
        for _ in tqdm.tqdm(pool.imap(generate_pdf, range(num_pdfs)), total=num_pdfs):
            pass
