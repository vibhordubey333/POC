import self as self
from borb.pdf import Document, Paragraph
from borb.pdf import Page
from borb.pdf import PDF
from borb.pdf.canvas.layout.page_layout.multi_column_layout import SingleColumnLayout
from borb.pdf.canvas.layout.annotation.link_annotation import (
    LinkAnnotation,
    DestinationType,
)
from decimal import Decimal
from borb.pdf.canvas.layout.image.image import Image
from borb.pdf.canvas.layout.table.fixed_column_width_table import FixedColumnWidthTable as Table
# --------------------Main method begins


# Create document
from pdfServe import _build_invoice_information, _build_billing_and_shipping_information, \
    _build_itemized_description_table

pdf = Document()

# Add page
page = Page()
pdf.add_page(page)

page_layout = SingleColumnLayout(page)
page_layout.vertical_margin = page.get_page_info().get_height() * Decimal(0.02)

# page_layout.add(
#     Image(
#         "download.png",width=Decimal(128),
#         height=Decimal(128),
#     ))*/
# Invoice information table
page_layout.add(_build_invoice_information())

# Empty paragraph for spacing
page_layout.add(Paragraph(" "))

# Invoice information table
page_layout.add(_build_invoice_information())

# Empty paragraph for spacing
page_layout.add(Paragraph(" "))

# Billing and shipping information table
page_layout.add(_build_billing_and_shipping_information())

# Itemized description
page_layout.add(_build_itemized_description_table(self))

with open("output2.pdf", "wb") as pdf_file_handle:
    PDF.dumps(pdf_file_handle, pdf)
# Outline
pdf.add_outline("Your Invoice", 0, DestinationType.FIT, page_nr=0)

# Main method ends.

# --------------------------------- References

#https://stackabuse.com/creating-pdf-invoices-in-python-with-borb/
#https://github.com/jorisschellekens/borb-examples/blob/master/README.md